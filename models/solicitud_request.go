package models

import "encoding/json"

// SolicitudRequest DTO para recibir la solicitud
type SolicitudRequest struct {
	SolicitudId     int             `json:"SolicitudId,omitempty"`
	TerceroId       int             `json:"TerceroId"`
	TipoSolicitudId string          `json:"TipoSolicitudId,omitempty"`
	SabaticoId      *int            `json:"SabaticoId,omitempty"`
	Formulario      json.RawMessage `json:"Formulario,omitempty"`
}
