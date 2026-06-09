package models

type AprobarRechazarPlanTRabajoSabaticoequest struct {
	SabaticoId            int    `json:"SabaticoId"`
	Justificacion         string `json:"Justificacion"`
	EstadoSabatico        string `json:"EstadoSabatico"`
	EstadoSoporteSabatico string `json:"EstadoSoporteSabatico"`
}
