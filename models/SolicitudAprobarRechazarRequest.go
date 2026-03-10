package models

type SolicitudAprobarRechazarRequest struct {
	TerceroId       int    `json:"TerceroId"`
	SolicitudId     int    `json:"SolicitudId"`
	Justificacion   string `json:"Justificacion"`
	EstadoSolicitud int    `json:"EstadoSolicitud"`
}
