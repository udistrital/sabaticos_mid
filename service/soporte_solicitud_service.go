package service

import (
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"
	"errors"
	"mime/multipart"
)

// CrearSoporteSolicitud procesa y guarda los documentos de soporte para una solicitud
func CrearSoporteSolicitud(soporteSolicitudReq models.SoporteSolicitudRequest, files []*multipart.FileHeader) (map[string]interface{}, error) {
	// Validar campos requeridos
	if soporteSolicitudReq.TerceroId == 0 || soporteSolicitudReq.SolicitudId == 0 ||
		soporteSolicitudReq.EstadoSolicitud == "" || soporteSolicitudReq.RolTercero == "" {
		return nil, errors.New("los campos terceroId, solicitudId, estadoSolicitud y rolTercero son requeridos")
	}

	// Procesar archivos si existen
	if len(files) > 0 {
		documentosGuardados, err := helpers.GuardarDocumentos(files, "soporte_solicitud")
		if err != nil {
			return nil, err
		}
		soporteSolicitudReq.Documentos = documentosGuardados
	}

	// Construir respuesta
	respuesta := map[string]interface{}{
		"ok":              true,
		"terceroId":       soporteSolicitudReq.TerceroId,
		"solicitudId":     soporteSolicitudReq.SolicitudId,
		"estadoSolicitud": soporteSolicitudReq.EstadoSolicitud,
		"rolTercero":      soporteSolicitudReq.RolTercero,
		"documentos":      soporteSolicitudReq.Documentos,
		"cantidad":        len(soporteSolicitudReq.Documentos),
	}

	return respuesta, nil
}
