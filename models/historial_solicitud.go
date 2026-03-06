package models

type HistorialSolicitud struct {
	Id                int         `json:"Id"`
	TerceroId         int         `json:"TerceroId"`
	Justificacion     string      `json:"Justificacion"`
	Activo            bool        `json:"Activo"`
	FechaCreacion     string      `json:"FechaCreacion,omitempty"`
	FechaModificacion string      `json:"FechaModificacion,omitempty"`
	EstadoSolicitudId interface{} `json:"EstadoSolicitudId,omitempty"`
	SolicitudId       interface{} `json:"SolicitudId,omitempty"`
}
