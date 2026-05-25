package models

type SoporteSabatcioRequest struct {
	SabaticoId            int    `json:"SabaticoId"`
	RolUsuario            string `json:"RolUsuario"`
	EstadoSoporteSabatico string `json:"EstadoSoporteSabatico"`
}
