package models

// SolicitudCreateRequest estructura para crear una solicitud
type SolicitudCreateRequest struct {
	TerceroId       int          `json:"TerceroId"`
	Activo          bool         `json:"Activo"`
	TipoSolicitudId IdReference  `json:"TipoSolicitudId"`
	SabaticoId      *IdReference `json:"SabaticoId,omitempty"`
}
