package models

// GestorDocumentalCreateRequest estructura para crear un gestor documental
type GestorDocumentalCreateRequest struct {
	IdTipoDocumento int         `json:"IdTipoDocumento"`
	Nombre          string      `json:"nombre"`
	Descripcion     string      `json:"descripcion,omitempty"`
	Metadatos       interface{} `json:"metadatos,omitempty"`
	File            string      `json:"file"`
}
