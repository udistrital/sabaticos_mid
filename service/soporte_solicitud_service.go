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

// CrearSoporteSolicitud procesa y guarda un documento de soporte para una solicitud
func CrearSoporteSolicitud(soporteSolicitudReq models.SoporteSolicitudRequest, file *multipart.FileHeader) (*models.SoporteSolicitudResponse, error) {
	descripcion := "Soporte para solicitud ID " + strconv.Itoa(soporteSolicitudReq.SolicitudId)

	// Convertir archivo a base64
	archivosBase64, err := helpers.ConvertirArchivosABase64([]*multipart.FileHeader{file})
	if err != nil {
		return nil, fmt.Errorf("error converting file to base64: %w", err)
	}

	//consultar tipo de documento para soporte de solicitud
	tipoDocumento, err := clients.ConsultarTipoDocumento(string(enums.SOLCITUD_SABATICO))
	if err != nil {
		return nil, fmt.Errorf("error querying document type: %w", err)
	}

	estadoSoposorteSolicitud, err := clients.ConsultarEstadoSoporteSolicitud(soporteSolicitudReq.EstadoSoporteSolicitud)
	if err != nil {
		return nil, fmt.Errorf("error querying support request status: %w", err)
	}

	if len(archivosBase64) == 0 {
		return nil, fmt.Errorf("no file content received")
	}

	archivo := archivosBase64[0]

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
		return nil, fmt.Errorf("error registering document in gestor documental (%s): %w", archivo.Nombre, err)
	}

	_, err = clients.RegistrarSoporteSolicitud(
		gestorGuardado.Id,
		soporteSolicitudReq.TerceroId,
		soporteSolicitudReq.SolicitudId,
		estadoSoposorteSolicitud.Id,
		soporteSolicitudReq.RolUsuario,
		soporteSolicitudReq.TipoDocumentoId,
	)
	if err != nil {
		return nil, fmt.Errorf("error registering support request for document %d: %w", gestorGuardado.Id, err)
	}

	// Construir respuesta
	respuesta := &models.SoporteSolicitudResponse{
		Ok:          true,
		TerceroId:   soporteSolicitudReq.TerceroId,
		SolicitudId: soporteSolicitudReq.SolicitudId,
		RolUsuario:  soporteSolicitudReq.RolUsuario,
		Documentos:  []*models.GestorDocumental{gestorGuardado},
	}

	return respuesta, nil
}
