package models

type EstadoSolicitud struct {
	Id                int    `json:"Id"`
	CodigoAbreviacion string `json:"CodigoAbreviacion"`
	Estado            string `json:"Estado"`
	Descripcion       string `json:"Descripcion"`
	Activo            bool   `json:"Activo"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
}
