package models

type TipoSolicitud struct {
	Id                int    `json:"Id"`
	CodigoAbreviacion string `json:"CodigoAbreviacion"`
	TipoSolicitud     string `json:"TipoSolicitud"`
	Descripcion       string `json:"Descripcion"`
	Activo            bool   `json:"Activo"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
}
