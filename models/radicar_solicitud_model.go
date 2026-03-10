package models

import "encoding/json"

type RadicarSolicitudRequest struct {
	SolicitudId  int             `json:"SolicitudId"`
	DocumentosId []int           `json:"DocumentosId,omitempty"`
	FormularioId int             `json:"FormularioId,omitempty"`
	Formulario   json.RawMessage `json:"Formulario,omitempty"`
}
