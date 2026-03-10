package models

type SolicitudAprobarRechazarResponse struct {
	SolicitudId     int `json:"SolicitudId"`
	EstadoSolicitud int `json:"EstadoSolicitud"`
}
