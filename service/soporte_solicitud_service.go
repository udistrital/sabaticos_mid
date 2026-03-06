package service

import (
	"api_mid_sabaticos/clients"
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"
	"fmt"
	"mime/multipart"
	"strconv"
)

// CrearSoporteSolicitud procesa y guarda los documentos de soporte para una solicitud
func CrearSoporteSolicitud(soporteSolicitudReq models.SoporteSolicitudRequest, files []*multipart.FileHeader) (*models.SoporteSolicitudResponse, error) {
	descripcion := "Soporte para solicitud ID " + strconv.Itoa(soporteSolicitudReq.SolicitudId)

	// Convertir archivos a base64
	archivosBase64, err := helpers.ConvertirArchivosABase64(files)
	if err != nil {
		return nil, fmt.Errorf("error convirtiendo archivos a base64: %w", err)
	}

	//consultar tipo de documento para soporte de solicitud
	tipoDocumento, err := clients.ConsultarTipoDocumento(string(enums.SOLCITUD_SABATICO))
	if err != nil {
		return nil, fmt.Errorf("error consultando tipo de documento: %w", err)
	}

	estadoSoposorteSolicitud, err := clients.ConsultarEstadoSoporteSolicitud(soporteSolicitudReq.EstadoSoporteSolicitud)
	if err != nil {
		return nil, fmt.Errorf("error consultando estado de soporte de solicitud: %w", err)
	}

	// Guardar cada documento en el gestor documental
	var documentosGuardados []*models.GestorDocumental

	for _, archivo := range archivosBase64 {
		// Construir metadatos específicos para el gestor documental
		metadatosGestor := map[string]interface{}{
			"NombreArchivo": archivo.Nombre,
			"Tipo":          "Archivo",
			"IdNuxeo":       "", // Será generado por el gestor documental
			"Observaciones": "Soporte de solicitud ID " + strconv.Itoa(soporteSolicitudReq.SolicitudId),
		}

		gestorGuardado, err := clients.RegistrarGestorDocumental(
			tipoDocumento.Id,
			archivo.Nombre,
			descripcion,
			metadatosGestor,
			archivo.Contenido,
		)
		if err != nil {
			return nil, fmt.Errorf("error registrando documento en gestor documental (%s): %w", archivo.Nombre, err)
		}
		documentosGuardados = append(documentosGuardados, gestorGuardado)

	}

	var soportesSoplicitud []models.SoporteSolicitud

	for _, documento := range documentosGuardados {
		soporteSolucitud, err := clients.RegistrarSoporteSolicitud(documento.Id, soporteSolicitudReq.TerceroId, soporteSolicitudReq.SolicitudId, estadoSoposorteSolicitud.Id, soporteSolicitudReq.RolUsuario)
		if err != nil {
			return nil, fmt.Errorf("error registrando soporte de solicitud para documento %d: %w", documento.Id, err)
		}

		soportesSoplicitud = append(soportesSoplicitud, *soporteSolucitud)
	}

	// Construir respuesta
	respuesta := &models.SoporteSolicitudResponse{
		Ok:                 true,
		CantidadDocumentos: len(soportesSoplicitud),
		TerceroId:          soporteSolicitudReq.TerceroId,
		SolicitudId:        soporteSolicitudReq.SolicitudId,
		RolUsuario:         soporteSolicitudReq.RolUsuario,
		Documentos:         documentosGuardados,
	}

	return respuesta, nil
}
