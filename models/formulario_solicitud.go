package models

type FormularioSolicitud struct {
	Id                int         `json:"Id"`
	Contenido         string      `json:"Contenido"`
	Activo            bool        `json:"Activo"`
	FechaCreacion     string      `json:"FechaCreacion,omitempty"`
	FechaModificacion string      `json:"FechaModificacion,omitempty"`
	SolicitudId       interface{} `json:"SolicitudId,omitempty"`
}
