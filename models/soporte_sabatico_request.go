package models

type SoporteSabatcioRequest struct {
	SabaticoId            int    `json:"SabaticoId"`
	RolUsuario            string `json:"RolUsuario"`
	NombreArchivo         string `json:"NombreArchivo"`
	EstadoSoporteSabatico string `json:"EstadoSoporteSabatico"`
}
