package models

// TipoTercero model
type TipoTercero struct {
	Id                int    `json:"Id"`
	Nombre            string `json:"Nombre"`
	Descripcion       string `json:"Descripcion,omitempty"`
	CodigoAbreviacion string `json:"CodigoAbreviacion,omitempty"`
	Activo            bool   `json:"Activo"`
	FechaCreacion     string `json:"FechaCreacion,omitempty"`
	FechaModificacion string `json:"FechaModificacion,omitempty"`
}
