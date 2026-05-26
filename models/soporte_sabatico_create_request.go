package models

type SoporteSabaticoCreateRequest struct {
	Activo                  bool        `json:"Activo"`
	FechaCreacion           string      `json:"FechaCreacion"`
	FechaModificacion       string      `json:"FechaModificacion"`
	DocumentoId             int         `json:"DocumentoId"`
	SabaticoId              IdReference `json:"SabaticoId"`
	RolUsuario              string      `json:"RolUsuario"`
	EstadoSoporteSabaticoId IdReference `json:"EstadoSoporteSabaticoId"`
}
