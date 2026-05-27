package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/enums"
	"github.com/udistrital/sabaticos_mid/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func CrearSolicitud(solicitudReq models.SolicitudRequest) (*models.Solicitud, error) {
	terceroId := solicitudReq.TerceroId
	codigoTipoSolicitud := solicitudReq.TipoSolicitudId
	sabaticoId := solicitudReq.SabaticoId
	formulario := solicitudReq.Formulario

	// Validar tercero
	//Se comenta por que por el momento no hay usuario de prueba en el servicio de terceros
	// if err := clients.ValidarTercero(terceroId); err != nil {
	// 	return nil, err
	// }

	tipoSolicitud, err := clients.ConsultarTipoSolicitud(codigoTipoSolicitud)
	if err != nil {
		return nil, err
	}

	if err := validarSolicitudPorTipo(tipoSolicitud.CodigoAbreviacion, sabaticoId); err != nil {
		return nil, err
	}

	// Crear solicitud en CRUD y obtener ID
	solicitud, err := clients.RegistrarSolicitud(terceroId, tipoSolicitud.Id, sabaticoId)
	if err != nil {
		return nil, err
	}

	// Determinar si debe crear formulario según el tipo de solicitud.
	// NUEVA, SUSPENSION y MODIFICACION requieren formulario asociado.
	debeCrearFormulario := tipoSolicitud.CodigoAbreviacion == string(enums.NUEVA) ||
		tipoSolicitud.CodigoAbreviacion == string(enums.SUSPENSION) ||
		tipoSolicitud.CodigoAbreviacion == string(enums.MODIFICACION)

	// Determinar el estado inicial según el tipo de solicitud:
	// NUEVA nace en BORRADOR; SUSPENSION y MODIFICACION se radican
	// automáticamente y nacen en RADICADA_ENVIADA_SA.
	estadoInicial := enums.BORRADOR
	if tipoSolicitud.CodigoAbreviacion == string(enums.SUSPENSION) ||
		tipoSolicitud.CodigoAbreviacion == string(enums.MODIFICACION) {
		estadoInicial = enums.RADICADA_ENVIADA_SA
	}

	_, _, err = registrarHistorialYFormulario(solicitud.Id, terceroId, string(formulario), string(estadoInicial), debeCrearFormulario)
	if err != nil {
		return nil, err
	}

	return solicitud, nil
}

func validarSolicitudPorTipo(CodigoAbreviacion string, sabaticoId *int) error {
	if CodigoAbreviacion == string(enums.NUEVA) {
		if sabaticoId != nil {
			return errors.New("a NEW request cannot be created with an associated Sabbatical")
		}
		return nil
	}

	switch CodigoAbreviacion {
	case string(enums.SUSPENSION):
		return validarSolicitudConSabatico(sabaticoId, "SUSPENSION")
	case string(enums.MODIFICACION):
		return validarSolicitudConSabatico(sabaticoId, "MODIFICATION")
	default:
		return nil
	}
}

// validarSolicitudConSabatico aplica las reglas comunes a tipos de solicitud
// que requieren un sabático asociado existente y dentro de los 3 meses desde
// su creación (actualmente SUSPENSION y MODIFICACION). El argumento tipoLabel
// se usa únicamente para construir mensajes de error claros para el cliente.
func validarSolicitudConSabatico(sabaticoId *int, tipoLabel string) error {
	if sabaticoId == nil {
		return fmt.Errorf("a %s request must have an associated Sabbatical", tipoLabel)
	}

	sabatico, err := clients.ConsultarSabatico(*sabaticoId)
	if err != nil {
		return err
	}

	// Validar que el sabático tenga máximo 3 meses desde su creación
	fechaCreacion, err := time.Parse("2006-01-02 15:04:05 -0700 -0700", sabatico.FechaCreacion)
	if err != nil {
		return errors.New("invalid FechaCreacion format for the Sabbatical")
	}

	fechaLimite := fechaCreacion.AddDate(0, 3, 0)
	if time.Now().After(fechaLimite) {
		return fmt.Errorf("a %s request cannot be created after 3 months from the Sabbatical creation date", tipoLabel)
	}

	return nil
}

func registrarHistorialYFormulario(solicitudId int, terceroId int, formularioRequest string, codigoEstadoSolicitud string, crearFormulario bool) (*models.HistorialSolicitud, *models.FormularioSolicitud, error) {
	var historial *models.HistorialSolicitud
	var formulario *models.FormularioSolicitud
	var historialErr, formularioErr error

	// Canal para sincronizar goroutines
	done := make(chan bool, 2)

	justificacion := "Creación de solicitud con estado " + codigoEstadoSolicitud

	// Crear historial en goroutine
	go func() {
		historial, historialErr = clients.RegistrarHistorialSolicitud(solicitudId, terceroId, justificacion, codigoEstadoSolicitud)
		if historialErr != nil {
			beego.Error("error registering request history:", historialErr)
		}
		done <- true
	}()

	// Solo crear formulario si el flag está habilitado
	// La validación de tipo de solicitud ya fue realizada en validarSolicitudPorTipo
	if crearFormulario {
		go func() {
			formulario, formularioErr = clients.RegistrarFormularioSolicitud(solicitudId, formularioRequest)
			if formularioErr != nil {
				beego.Error("error registering request form:", formularioErr)
			}
			done <- true
		}()
	} else {
		// Si no se crea formulario, marcar como completado
		done <- true
	}

	// Esperar a que ambas goroutines terminen
	<-done
	<-done

	// Verificar errores
	if historialErr != nil {
		return nil, nil, historialErr
	}
	if formularioErr != nil {
		return nil, nil, formularioErr
	}

	return historial, formulario, nil
}

func CambiarEstado(SolicitudAprobarRechazarRequest models.SolicitudAprobarRechazarRequest) (*models.HistorialSolicitud, error) {

	IdEstado, err := clients.ConsultarEstadoSolicitud(SolicitudAprobarRechazarRequest.EstadoSolicitud)
	if err != nil {
		return nil, err
	}

	if err := desactivarRegistrosHistorial(SolicitudAprobarRechazarRequest.SolicitudId); err != nil {
		return nil, err
	}

	/*
		Si la solicitud es aprobar y enviar la solicitud,
		se aprueban todos los documentos asociados
	*/

	if SolicitudAprobarRechazarRequest.EstadoSoporte != "" {

		// Obtener soportes asociados a la solicitud
		soportes, err := clients.ConsultarSoportesSolicitud(SolicitudAprobarRechazarRequest.SolicitudId)
		if err != nil {
			logs.Error("Error consultando soportes:", err)
			return nil, err
		}

		// Validar si hay soportes
		if len(soportes) == 0 {
			return nil, errors.New("no soportes found for solicitud: " + fmt.Sprint(SolicitudAprobarRechazarRequest.SolicitudId))
		}

		for _, soporte := range soportes {

			_, err := clients.ActualizarSoporteSolicitud(
				soporte.DocumentoId,
				SolicitudAprobarRechazarRequest.SolicitudId,
				SolicitudAprobarRechazarRequest.EstadoSoporte,
			)

			if err != nil {
				logs.Error("Error actualizando soporte:", soporte.Id, err)
				continue
			}
		}

	}

	HistorialSolicitudEstado, err := clients.RegistrarHistorialSolicitudEstado(
		SolicitudAprobarRechazarRequest.SolicitudId,
		SolicitudAprobarRechazarRequest.TerceroId,
		SolicitudAprobarRechazarRequest.Justificacion,
		IdEstado.Id,
	)
	if err != nil {
		return nil, err
	}

	return HistorialSolicitudEstado, err
}

func desactivarRegistrosHistorial(solicitudId int) error {
	idsHistorial, err := clients.ConsultarIdsHistorialSolicitud(solicitudId)
	if err != nil {
		return err
	}

	for _, idHistorial := range idsHistorial {
		_, err := clients.DesactivarHistorialSolicitud(idHistorial)
		if err != nil {
			return err
		}
	}

	return nil
}

func extraerFacultadFormulario(formulario models.FormularioSolicitud) (string, error) {
	var contenidoJSON struct {
		Docente struct {
			Facultad string `json:"facultad"`
		} `json:"docente"`
	}

	if err := json.Unmarshal([]byte(formulario.Contenido), &contenidoJSON); err != nil {
		return "", fmt.Errorf("invalid json in formulario %d: %w", formulario.Id, err)
	}

	return strings.TrimSpace(contenidoJSON.Docente.Facultad), nil
}

func consultarHistorialesPorSolicitud(solicitudId int, estadosSolicitud []string) ([]models.HistorialSolicitud, error) {
	idsHistorial, err := clients.ConsultarIdsHistorialSolicitud(solicitudId)
	if err != nil {
		return nil, err
	}

	historialSolicitud := make([]models.HistorialSolicitud, 0)
	for _, idHistorial := range idsHistorial {
		historial, err := clients.ConsultarHistorialSolicitudIdEstadoId(idHistorial, estadosSolicitud)
		if err != nil {
			return nil, err
		}

		historialSolicitud = append(historialSolicitud, historial...)
	}

	return historialSolicitud, nil
}

func GetFormulariosByDocumentoId(documentoId string, estadosSolicitud []string) ([]models.HistorialSolicitud, error) {
	SecretariaAcademicaUsuario, err := clients.ConsultarSecretariaAcademicaDocumentoUserId(documentoId)
	if err != nil {
		return nil, err
	}

	formularios, err := clients.ConsultarTodosFormulariosSolicitud()
	if err != nil {
		return nil, err
	}

	dependenciaUsuario := strings.TrimSpace(SecretariaAcademicaUsuario.Dependencia)
	if dependenciaUsuario == "" {
		return nil, errors.New("dependencia de secretaria academica vacía")
	}

	historialSolicitud := make([]models.HistorialSolicitud, 0)
	solicitudesProcesadas := make(map[int]bool)

	for _, formulario := range formularios {
		facultadFormulario, err := extraerFacultadFormulario(formulario)
		if err != nil {
			return nil, err
		}

		if !strings.EqualFold(facultadFormulario, dependenciaUsuario) {
			continue
		}

		solicitudId := formulario.SolicitudId.Id
		if solicitudesProcesadas[solicitudId] {
			continue
		}

		solicitudesProcesadas[solicitudId] = true

		historiales, err := consultarHistorialesPorSolicitud(solicitudId, estadosSolicitud)
		if err != nil {
			return nil, err
		}

		historialSolicitud = append(historialSolicitud, historiales...)
	}

	return historialSolicitud, nil
}

func RadicarSolicitud(RadicarSolicitudRequest models.RadicarSolicitudRequest) (map[string]interface{}, error) {

	solicitud, err := clients.ConsultarSolicitud(RadicarSolicitudRequest.SolicitudId)
	if err != nil {
		return nil, err
	}

	err = desactivarRegistrosHistorial(RadicarSolicitudRequest.SolicitudId)
	if err != nil {
		return nil, err
	}

	justificacion := "Radicación de solicitud"

	historialSolicitud, err := clients.RegistrarHistorialSolicitud(solicitud.Id, solicitud.TerceroId, justificacion, string(enums.RADICADA_ENVIADA_SA))
	if err != nil {
		beego.Error("error registering request history:", err)
	}
	formularioActualizado, err := clients.ActualizarFormularioSolicitud(solicitud.Id, RadicarSolicitudRequest.FormularioId, string(RadicarSolicitudRequest.Formulario))
	if err != nil {
		return nil, err
	}

	var soportes []*models.SoporteSolicitud
	for _, soporteId := range RadicarSolicitudRequest.DocumentosId {
		soporte, err := clients.ActualizarSoporteSolicitud(soporteId, solicitud.Id, string(enums.SA_PENDIENTE_REVISION_SA)) // Validar que sea ese por el cambio de radicado
		if err != nil {
			return nil, err
		}
		soportes = append(soportes, soporte)
	}

	response := map[string]interface{}{
		"solicitud":  solicitud,
		"historial":  historialSolicitud,
		"formulario": formularioActualizado,
		"soportes":   soportes,
	}

	return response, nil
}
