package models

type ValidarSolicitudRequest struct {
	TerceroId         int    `json:"tercero_id"`
	SolicitudId       int    `json:"solicitud_id"`
	EstadoSolicitudId int    `json:"estado_solicitud_id,omitempty"`
	Justificacion     string `json:"justificacion,omitempty"`
}
