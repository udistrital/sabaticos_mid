package models

type CrearSabaticoCrudRequest struct {
	TerceroId         int                    `json:"TerceroId"`
	Observaciones     string                 `json:"Observaciones"`
	FechaInicio       string                 `json:"FechaInicio"`
	FechaFin          string                 `json:"FechaFin"`
	Activo            bool                   `json:"Activo"`
	FechaCreacion     string                 `json:"FechaCreacion"`
	FechaModificacion string                 `json:"FechaModificacion"`
	EstadoSabaticoId  map[string]interface{} `json:"EstadoSabaticoId"`
}
