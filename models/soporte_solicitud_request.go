package models

// SoporteSolicitudRequest estructura para recibir soporte con múltiples documentos
type SoporteSolicitudRequest struct {
	TerceroId              int      `json:"tercero_id"`
	TipoDocumentoId        int      `json:"tipo_documento_id"`
	EstadoSoporteSolicitud string   `json:"estado_soporte_solicitud_id"`
	SolicitudId            int      `json:"solicitud_id"`
	RolUsuario             string   `json:"rol_tercero"`
	Documentos             []string `json:"documentos"` // Base64 o rutas de archivos
}
