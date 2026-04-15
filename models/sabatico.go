package models

// Sabatico modelo de un sabático
type Sabatico struct {
	Id                int         `json:"Id"`
	TerceroId         int         `json:"TerceroId"`
	EstadoSabaticoId  interface{} `json:"EstadoSabaticoId"`
	FechaInicio       string      `json:"FechaInicio"`
	FechaFin          string      `json:"FechaFin"`
	Observaciones     string      `json:"Observaciones"`
	Activo            bool        `json:"Activo"`
	FechaCreacion     string      `json:"FechaCreacion"`
	FechaModificacion string      `json:"FechaModificacion"`
}
