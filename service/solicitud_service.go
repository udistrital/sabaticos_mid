package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/enums"
	"github.com/udistrital/sabaticos_mid/models"

	"github.com/astaxie/beego"
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

	// Determinar si debe crear formulario según el tipo de solicitud
	debeCrearFormulario := tipoSolicitud.CodigoAbreviacion == string(enums.NUEVA)

	// Crear historial y formulario en paralelo
	_, _, err = registrarHistorialYFormulario(solicitud.Id, terceroId, string(formulario), string(enums.BORRADOR), debeCrearFormulario)
	if err != nil {
		return nil, err
	}

	// Solo retornar la solicitud si TODO fue exitoso
	return solicitud, nil
}

func validarSolicitudPorTipo(CodigoAbreviacion string, sabaticoId *int) error {
	if CodigoAbreviacion == string(enums.NUEVA) {
		if sabaticoId != nil {
			return errors.New("a NEW request cannot be created with an associated Sabbatical")
		}
		return nil
	}

	if CodigoAbreviacion != string(enums.SUSPENSION) {
		return nil
	}

	if sabaticoId == nil {
		return errors.New("a SUSPENSION request must have an associated Sabbatical")
	}

	// Consultar si el sabático existe
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
		return errors.New("a SUSPENSION request cannot be created after 3 months from the Sabbatical creation date")
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
		fmt.Printf("Error consultando tipo de solicitud: %v\n", err)
		return nil, err
	}

	/* Se consultan los ids de la tabla historial asociados a la solicitud. */
	idsHistorial, err := clients.ConsultarIdsHistorialSolicitud(SolicitudAprobarRechazarRequest.SolicitudId)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	if len(idsHistorial) > 0 {
		/* Se desactivan los historiales asociados */
		for _, idHistorial := range idsHistorial {
			_, err := clients.DesactivarHistorialSolicitud(idHistorial)
			if err != nil {
				fmt.Println("error desactivando historial:", err)
				return nil, err
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

	return HistorialSolicitudEstado, nil
}

func Rechazar(SolicitudRechazarRequest models.SolicitudAprobarRechazarRequest) (*models.HistorialSolicitud, error) {
	return nil, nil
}

func RadicarSolicitud(RadicarSolicitudRequest models.RadicarSolicitudRequest) (map[string]interface{}, error) {

	solicitud, err := clients.ConsultarSolicitud(RadicarSolicitudRequest.SolicitudId)
	if err != nil {
		return nil, err
	}

	justificacion := "Radicación de solicitud"

	historialSolicitud, err := clients.RegistrarHistorialSolicitud(solicitud.Id, solicitud.TerceroId, justificacion, string(enums.RADICADA_ENVIADA_SA_RADICADA))
	if err != nil {
		beego.Error("error registering request history:", err)
	}

	formularioActtualizado, err := clients.ActualizarFormularioSolicitud(solicitud.Id, RadicarSolicitudRequest.FormularioId, string(RadicarSolicitudRequest.Formulario))
	if err != nil {
		return nil, err
	}

	var soportes []*models.SoporteSolicitud
	for _, soporteId := range RadicarSolicitudRequest.DocumentosId {
		soporte, err := clients.ActualizarSoporteSolicitud(soporteId, solicitud.Id, string(enums.RADICADO))
		if err != nil {
			return nil, err
		}
		soportes = append(soportes, soporte)
	}

	response := map[string]interface{}{
		"solicitud":  solicitud,
		"historial":  historialSolicitud,
		"formulario": formularioActtualizado,
		"soportes":   soportes,
	}

	return response, nil
}
