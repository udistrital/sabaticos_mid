package models

type SolicitudAprobarRechazarRequest struct {
	TerceroId       int    `json:"TerceroId"`
	SolicitudId     int    `json:"SolicitudId"`
	Justificacion   string `json:"Justificacion"`
	EstadoSolicitud string `json:"EstadoSolicitud"`
	EstadoSoporte   string
}
