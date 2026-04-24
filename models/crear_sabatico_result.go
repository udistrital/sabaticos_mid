package models

type CrearSabaticoResult struct {
	Id                int         `json:"Id"`
	TerceroId         int         `json:"TerceroId"`
	Observaciones     string      `json:"Observaciones"`
	FechaInicio       string      `json:"FechaInicio"`
	FechaFin          string      `json:"FechaFin"`
	Activo            bool        `json:"Activo"`
	FechaCreacion     string      `json:"FechaCreacion"`
	FechaModificacion string      `json:"FechaModificacion"`
	EstadoSabaticoId  interface{} `json:"EstadoSabaticoId"`
}
