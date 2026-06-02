package models

// Tercero model minimal para referencias
type Tercero struct {
	Id             int    `json:"Id"`
	NombreCompleto string `json:"NombreCompleto,omitempty"`
}
