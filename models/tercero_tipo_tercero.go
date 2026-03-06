package models

// TerceroTipoTercero model usado en responses desde el CRUD
type TerceroTipoTercero struct {
	Id                int         `json:"Id"`
	TerceroId         interface{} `json:"TerceroId,omitempty"`
	TipoTerceroId     interface{} `json:"TipoTerceroId,omitempty"`
	Activo            bool        `json:"Activo"`
	FechaCreacion     string      `json:"FechaCreacion,omitempty"`
	FechaModificacion string      `json:"FechaModificacion,omitempty"`
}
