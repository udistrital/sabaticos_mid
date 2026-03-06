package models

// SoporteSolicitudResponse estructura de respuesta para la creación de soportes de solicitud
type SoporteSolicitudResponse struct {
	Ok                 bool                `json:"ok"`
	CantidadDocumentos int                 `json:"cantidadDocumentos"`
	TerceroId          int                 `json:"terceroId"`
	SolicitudId        int                 `json:"solicitudId"`
	RolUsuario         string              `json:"rolUsuario"`
	Documentos         []*GestorDocumental `json:"documentos"`
}
