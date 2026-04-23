package models

// SoporteSolicitud modelo de la base de datos adaptado para el mid
type SoporteSolicitud struct {
	Id                       int         `json:"Id"`
	DocumentoId              int         `json:"DocumentoId"`
	TerceroId                int         `json:"TerceroId"`
	Activo                   bool        `json:"Activo"`
	FechaCreacion            string      `json:"FechaCreacion,omitempty"`
	FechaModificacion        string      `json:"FechaModificacion,omitempty"`
	SolicitudId              interface{} `json:"SolicitudId,omitempty"`
	EstadoSoporteSolicitudId interface{} `json:"EstadoSoporteSolicitudId,omitempty"`
	RolUsuario               string      `json:"RolUsuario,omitempty"`
	TipoDocumentoId          int         `json:"TipoDocumentoId,omitempty"`
}
