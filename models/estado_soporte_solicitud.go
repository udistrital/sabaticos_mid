package models

type EstadoSoporteSolicitud struct {
	Id                int    `json:"Id"`
	CodigoAbreviacion string `json:"CodigoAbreviacion"`
	NombreEstado      string `json:"NombreEstado"`
	Activo            bool   `json:"Activo"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
}
