package models

// TipoDocumento model para el tipo de documento
type TipoDocumento struct {
	Id                   int         `json:"Id"`
	Nombre               string      `json:"Nombre"`
	Descripcion          string      `json:"Descripcion,omitempty"`
	CodigoAbreviacion    string      `json:"CodigoAbreviacion,omitempty"`
	Activo               bool        `json:"Activo"`
	NumeroOrden          int         `json:"NumeroOrden,omitempty"`
	Tamano               int         `json:"Tamano,omitempty"`
	Extension            string      `json:"Extension,omitempty"`
	Workspace            string      `json:"Workspace,omitempty"`
	TipoDocumentoNuxeo   string      `json:"TipoDocumentoNuxeo,omitempty"`
	FechaCreacion        string      `json:"FechaCreacion,omitempty"`
	FechaModificacion    string      `json:"FechaModificacion,omitempty"`
	DominioTipoDocumento interface{} `json:"DominioTipoDocumento,omitempty"`
}
