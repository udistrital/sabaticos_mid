package models

// GestorDocumental model para el gestor documental completo
type GestorDocumental struct {
	Id                int         `json:"Id"`
	Nombre            string      `json:"Nombre"`
	Descripcion       string      `json:"Descripcion,omitempty"`
	Enlace            string      `json:"Enlace,omitempty"`
	TipoDocumento     interface{} `json:"TipoDocumento,omitempty"`
	Metadatos         string      `json:"Metadatos,omitempty"`
	Activo            bool        `json:"Activo"`
	FechaCreacion     string      `json:"FechaCreacion,omitempty"`
	FechaModificacion string      `json:"FechaModificacion,omitempty"`
}
