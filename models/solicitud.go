package models

// Solicitud modelo de la base de datos
type Solicitud struct {
	Id                int         `json:"Id"`
	TerceroId         int         `json:"TerceroId"`
	Activo            bool        `json:"Activo"`
	FechaCreacion     string      `json:"FechaCreacion,omitempty"`
	FechaModificacion string      `json:"FechaModificacion,omitempty"`
	TipoSolicitudId   interface{} `json:"TipoSolicitudId,omitempty"`
	SabaticoId        interface{} `json:"SabaticoId,omitempty"`
}
