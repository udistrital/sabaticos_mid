package models

type HistorialEstadoSabatico struct {
	Id                int            `json:"Id"`
	TerceroId         int            `json:"TerceroId"`
	Justificacion     string         `json:"Justificacion"`
	Activo            bool           `json:"Activo"`
	FechaCreacion     string         `json:"FechaCreacion"`
	FechaModificacion string         `json:"FechaModificacion"`
	EstadoSabaticoId  EstadoSabatico `json:"EstadoSabaticoId"`
	SabaticoId        Sabatico       `json:"SabaticoId"`
}
