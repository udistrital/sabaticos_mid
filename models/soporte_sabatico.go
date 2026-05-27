package models

type SoporteSabatico struct {
	DocumentoId             int         `json:"DocumentoId"`
	Activo                  bool        `json:"Activo"`
	FechaCreacion           string      `json:"FechaCreacion"`
	FechaModificacion       string      `json:"FechaModificacion"`
	SabaticoId              IdReference `json:"SabaticoId"`
	RolUsuario              string      `json:"RolUsuario"`
	EstadoSoporteSabaticoId IdReference `json:"EstadoSoporteSabaticoId"`
	Id                      int         `json:"Id"`
}
