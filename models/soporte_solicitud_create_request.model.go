package models

// SoporteSolicitudCreateRequest estructura para crear un soporte de solicitud
type SoporteSolicitudCreateRequest struct {
	DocumentoId              int         `json:"DocumentoId"`
	TerceroId                int         `json:"TerceroId"`
	Activo                   bool        `json:"Activo"`
	SolicitudId              IdReference `json:"SolicitudId"`
	EstadoSoporteSolicitudId IdReference `json:"EstadoSoporteSolicitudId"`
	RolUsuario               string      `json:"RolUsuario"`
}
