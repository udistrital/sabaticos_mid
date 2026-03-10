package models

type SolicitudAprobarRechazarRequest struct {
	TerceroId       int    `json:"TerceroId"`
	SolicitudId     int    `json:"SolicitudId"`
	Justificacion   string `json:"Justificacion,omitempty"`
	EstadoSolicitud int    `json:"EstadoSolicitud"`
}
