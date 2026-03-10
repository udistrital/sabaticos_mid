package models

type SolicitudAprobarResponse struct {
	SolicitudId     int `json:"SolicitudId"`
	EstadoSolicitud int `json:"EstadoSolicitud"`
}
