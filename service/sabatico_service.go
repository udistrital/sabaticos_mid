package service

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/enums"
	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"
)

func CrearSabatico(
	solicitudId int,
	terceroId int,
	observaciones string,
	fechaInicio string,
	fechaFin string,
) (*models.CrearSabaticoResult, error) {

	estadoRequest := models.SolicitudAprobarRechazarRequest{
		TerceroId:       terceroId,
		SolicitudId:     solicitudId,
		Justificacion:   "Cambio automático al crear año sabático",
		EstadoSolicitud: "S12",
		EstadoSoporte:   "SGOK",
	}

	_, err := CambiarEstado(estadoRequest)
	if err != nil {
		return nil, fmt.Errorf(
			"error cambiando estado: %v",
			err,
		)
	}

	sabaticoCreado, err := clients.RegistrarSabatico(
		solicitudId,
		terceroId,
		observaciones,
		fechaInicio,
		fechaFin,
		"ES0",
	)

	if err != nil {
		return nil, err
	}

	err = clients.AsociarSabaticoSolicitud(
		solicitudId,
		sabaticoCreado.Id,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error asociando sabático a solicitud: %v",
			err,
		)
	}

	return sabaticoCreado, nil
}

func GuardarPlanTrabajoSabatico(
	request models.PlanTrabajoSabaticoRequest,
) (*models.HistorialEstadoSabatico, error) {

	historiales, err := clients.ConsultarHistorialEstadoSabatico(
		request.SabaticoId,
	)
	if err != nil {
		return nil, err
	}

	var planTrabajoExistente *models.HistorialEstadoSabatico

	// Buscar si existe ES1
	for _, historial := range historiales {
		if historial.EstadoSabaticoId.CodigoAbreviacion == "ES1" {
			historialCopy := historial
			planTrabajoExistente = &historialCopy
			break
		}
	}

	// Si existe ES1 -> editar
	if planTrabajoExistente != nil {

		planTrabajoExistente.Justificacion =
			request.Justificacion

		err = clients.ActualizarHistorialEstadoSabatico(
			*planTrabajoExistente,
		)
		if err != nil {
			return nil, err
		}

		return planTrabajoExistente, nil
	}

	// Si NO existe -> desactivar historiales
	for _, historial := range historiales {
		err := clients.DesactivarHistorialEstadoSabatico(
			historial,
		)
		if err != nil {
			return nil, err
		}
	}

	// Crear nuevo ES1
	estadoSabaticoId, err :=
		clients.ConsultarIdEstadoSabatico("ES1")
	if err != nil {
		return nil, err
	}

	planTrabajo, err :=
		clients.CrearHistorialEstadoSabatico(
			request.TerceroId,
			request.Justificacion,
			estadoSabaticoId,
			request.SabaticoId,
		)

	if err != nil {
		return nil, err
	}

	return planTrabajo, nil
}

func CambiarEstadoPlanTrabajoSabatico(cambiarEstadoPlanTrabajoRequest models.AprobarRechazarPlanTRabajoSabaticoequest) (*models.HistorialEstadoSabatico, error) {

	estadoSoporteSabatico, err := clients.ConsultarEstadoSoporteSabatico(cambiarEstadoPlanTrabajoRequest.EstadoSoporteSabatico)
	if err != nil {
		return nil, err
	}

	soportesSabatico, err := clients.ConsultarSoportesSabaticos(cambiarEstadoPlanTrabajoRequest.SabaticoId)
	if err != nil {
		return nil, err
	}

	historialesEstadoSabatico, err := clients.ConsultarHistorialEstadoSabatico(cambiarEstadoPlanTrabajoRequest.SabaticoId)
	if err != nil {
		return nil, err
	}

	estadoSabatico, err := clients.ConsultarIdEstadoSabatico(cambiarEstadoPlanTrabajoRequest.EstadoSabatico)
	if err != nil {
		return nil, err
	}

	for _, historial := range historialesEstadoSabatico {
		err := clients.DesactivarHistorialEstadoSabatico(historial)
		if err != nil {
			return nil, err
		}
	}

	for _, soporte := range soportesSabatico {
		soporte.EstadoSoporteSabaticoId = models.EstadoSabatico{Id: estadoSoporteSabatico.Id}
		_, err := clients.ActualizarSoporteSabatico(soporte)
		if err != nil {
			return nil, err
		}
	}

	if len(historialesEstadoSabatico) == 0 {
		return nil, fmt.Errorf("no historial available to obtain terceroId for sabatico %d", cambiarEstadoPlanTrabajoRequest.SabaticoId)
	}

	terceroIdFromHistorial := historialesEstadoSabatico[0].TerceroId

	historialEstadoSabatico, err := clients.CrearHistorialEstadoSabatico(
		terceroIdFromHistorial,
		cambiarEstadoPlanTrabajoRequest.Justificacion,
		estadoSabatico,
		cambiarEstadoPlanTrabajoRequest.SabaticoId,
	)

	if err != nil {
		return nil, err
	}

	return historialEstadoSabatico, nil
}

func ConsultarSoportesSabaticos(sabaticoId int) ([]map[string]interface{}, error) {
	var response []map[string]interface{}
	if sabaticoId <= 0 {
		return nil, fmt.Errorf("sabaticoId is required and must be a positive integer")
	}

	SoportesSabaticos, err := clients.ConsultarSoportesSabaticos(sabaticoId)
	if err != nil {
		return nil, err
	}

	for i := range SoportesSabaticos {
		gestorDocumental, err := clients.ConsultarGestorDocumental(SoportesSabaticos[i].DocumentoId)
		if err != nil {
			return nil, fmt.Errorf("error consulting gestor documental for support %d: %w", SoportesSabaticos[i].Id, err)
		}
		response = append(response, map[string]interface{}{
			"Id":                    SoportesSabaticos[i].Id,
			"SabaticoId":            SoportesSabaticos[i].SabaticoId,
			"EstadoSoporteSabatico": SoportesSabaticos[i].EstadoSoporteSabaticoId,
			"RolUsuario":            SoportesSabaticos[i].RolUsuario,
			"Documento":             gestorDocumental,
		})
	}

	return response, nil
}

func CrearSoporteSabatico(soporteSabaticoReq models.SoporteSabatcioRequest, file *multipart.FileHeader) (*models.SoporteSolicitudResponse, error) {
	descripcion := "Soporte para plan de trabajo para el Sabatico ID " + fmt.Sprint(soporteSabaticoReq.SabaticoId)

	archivosBase64, err := helpers.ConvertirArchivosABase64([]*multipart.FileHeader{file})
	if err != nil {
		return nil, fmt.Errorf("error converting file to base64: %w", err)
	}

	//consultar tipo de documento para soporte de solicitud
	tipoDocumento, err := clients.ConsultarTipoDocumento(string(enums.PLAN_TRABAJO))
	if err != nil {
		return nil, fmt.Errorf("error querying document type: %w", err)
	}

	estadoSoporteSabatico, err := clients.ConsultarEstadoSoporteSabatico(soporteSabaticoReq.EstadoSoporteSabatico)
	if err != nil {
		return nil, fmt.Errorf("error querying support request status: %w", err)
	}

	if len(archivosBase64) == 0 {
		return nil, fmt.Errorf("no file content received")
	}

	archivo := archivosBase64[0]

	metadatosGestor := map[string]interface{}{
		"NombreArchivo": archivo.Nombre,
		"Tipo":          "Archivo",
		"IdNuxeo":       "", // Será generado por el gestor documental
		"Observaciones": "Soporte de plan de trabajo para el Sabatico ID " + strconv.Itoa(soporteSabaticoReq.SabaticoId),
	}

	gestorGuardado, err := clients.RegistrarGestorDocumental(
		tipoDocumento.Id,
		soporteSabaticoReq.NombreArchivo,
		descripcion,
		metadatosGestor,
		archivo.Contenido,
	)

	if err != nil {
		return nil, fmt.Errorf("error registering document in gestor documental (%s): %w", archivo.Nombre, err)
	}

	_, err = clients.RegistrarSoporteSabatico(
		gestorGuardado.Id,
		soporteSabaticoReq.SabaticoId,
		estadoSoporteSabatico.Id,
		soporteSabaticoReq.RolUsuario,
	)
	if err != nil {
		return nil, fmt.Errorf("error registering support request for document %d: %w", gestorGuardado.Id, err)
	}

	// Construir respuesta
	respuesta := &models.SoporteSolicitudResponse{
		Ok: true,
		Documentos: []*models.GestorDocumental{
			gestorGuardado,
		},
		SolicitudId: soporteSabaticoReq.SabaticoId,
		RolUsuario:  soporteSabaticoReq.RolUsuario,
	}

	return respuesta, nil

}

func ConsultarSoportesSabaticosPorDocumentoSecretaria(documento string) ([]models.HistorialEstadoSabatico, error) {
	response := []models.HistorialEstadoSabatico{}
	sabaticosId := []int{}

	persona, err := clients.ConsultarSecretariaAcademicaDocumentoUserId(documento)
	if err != nil {
		return nil, fmt.Errorf(
			"error consulting secretaria academica for documento %s: %w",
			documento,
			err,
		)
	}

	formularios, err := clients.ConsultarFormularioTipoSolicitudSabatico(enums.NUEVA)
	if err != nil {
		return nil, fmt.Errorf(
			"error consulting forms for new sabatico request: %w",
			err,
		)
	}

	for _, formulario := range *formularios {
		var contenidoJSON map[string]interface{}

		if err := json.Unmarshal(
			[]byte(formulario.Contenido),
			&contenidoJSON,
		); err != nil {
			return nil, fmt.Errorf(
				"formulario %d: error parsing json content: %w",
				formulario.Id,
				err,
			)
		}

		// Extraer plan_trabajo_ano_sabatico
		planTrabajo, ok := contenidoJSON["plan_trabajo_ano_sabatico"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf(
				"formulario %d: missing or invalid plan_trabajo_ano_sabatico structure",
				formulario.Id,
			)
		}

		// Extraer identificacion_docente
		identificacionDocente, ok := planTrabajo["identificacion_docente"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf(
				"formulario %d: missing or invalid identificacion_docente structure",
				formulario.Id,
			)
		}

		// Extraer facultad
		facultadFormulario, ok := identificacionDocente["facultad"].(string)
		if !ok {
			return nil, fmt.Errorf(
				"formulario %d: missing or invalid facultad field",
				formulario.Id,
			)
		}

		// Comparar facultad
		if facultadFormulario == persona.Dependencia {
			sabaticosId = append(
				sabaticosId,
				formulario.SolicitudId.SabaticoId.Id,
			)
		}
	}

	// Consultar soportes
	for _, sabaticoId := range sabaticosId {
		historial, err := clients.ConsultarporEstadoHistorialEstadoSabatico(sabaticoId, enums.EstadoSabatico(enums.REVISION_SA))
		if err != nil {
			return nil, fmt.Errorf(
				"error consulting sabatico %d: %w",
				sabaticoId,
				err,
			)
		}

		response = append(response, historial...)
	}

	return response, nil
}
