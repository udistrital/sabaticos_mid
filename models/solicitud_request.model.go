package models

// SolicitudRequest DTO para recibir la solicitud
type SolicitudRequest struct {
	TerceroId       int         `json:"TerceroId"`
	TipoSolicitudId interface{} `json:"TipoSolicitudId,omitempty"`
	SabaticoId      interface{} `json:"SabaticoId,omitempty"`
}
