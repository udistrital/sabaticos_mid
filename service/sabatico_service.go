package service

import (
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

		estado, ok :=
			historial.EstadoSabaticoId.(map[string]interface{})
		if !ok {
			continue
		}

		codigo, ok :=
			estado["CodigoAbreviacion"].(string)
		if !ok {
			continue
		}

		if codigo == "ES1" {
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
		soporte.EstadoSoporteSabaticoId = models.IdReference{Id: estadoSoporteSabatico.Id}
		fmt.Println("hello")
		fmt.Println(soporte)
		_, err := clients.ActualizarSoporteSabatico(soporte)
		if err != nil {
			return nil, err
		}
	}

	historialEstadoSabatico, err := clients.CrearHistorialEstadoSabatico(
		cambiarEstadoPlanTrabajoRequest.TerceroId,
		cambiarEstadoPlanTrabajoRequest.Justificacion,
		estadoSabatico,
		cambiarEstadoPlanTrabajoRequest.SabaticoId,
	)

	if err != nil {
		return nil, err
	}

	return historialEstadoSabatico, nil
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
		archivo.Nombre,
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
