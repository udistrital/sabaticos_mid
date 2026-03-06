package models

// HistorialSolicitudCreateRequest estructura para crear un historial de solicitud
type HistorialSolicitudCreateRequest struct {
	TerceroId         int         `json:"TerceroId"`
	Justificacion     string      `json:"Justificacion"`
	Activo            bool        `json:"Activo"`
	EstadoSolicitudId IdReference `json:"EstadoSolicitudId"`
	SolicitudId       IdReference `json:"SolicitudId"`
}
