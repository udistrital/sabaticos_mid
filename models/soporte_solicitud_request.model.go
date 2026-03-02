package models

// SoporteSolicitudRequest estructura para recibir soporte con múltiples documentos
type SoporteSolicitudRequest struct {
	TerceroId       int      `json:"tercero_id"`
	SolicitudId     int      `json:"solicitud_id"`
	EstadoSolicitud string   `json:"estado_solicitud"`
	RolTercero      string   `json:"rol_tercero"`
	Documentos      []string `json:"documentos"` // Base64 o rutas de archivos
}
