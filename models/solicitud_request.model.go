package models

// SolicitudRequest DTO para recibir la solicitud
type SolicitudRequest struct {
	TerceroId       int    `json:"TerceroId"`
	TipoSolicitudId string `json:"TipoSolicitudId,omitempty"`
	SabaticoId      *int   `json:"SabaticoId,omitempty"`
}
