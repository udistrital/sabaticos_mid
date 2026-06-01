package models

type FormularioSolicitud struct {
	Id                int             `json:"Id"`
	Contenido         string          `json:"Contenido"`
	Activo            bool            `json:"Activo"`
	FechaCreacion     string          `json:"FechaCreacion,omitempty"`
	FechaModificacion string          `json:"FechaModificacion,omitempty"`
	SolicitudId       SolicitudMinima `json:"SolicitudId,omitempty"`
}

type SolicitudMinima struct {
	Id         int            `json:"Id"`
	SabaticoId SabaticoMinimo `json:"SabaticoId"`
}

type SabaticoMinimo struct {
	Id int `json:"Id"`
}
