package models

type PlanTrabajoSabaticoRequest struct {
	TerceroId     int    `json:"TerceroId"`
	Justificacion string `json:"Justificacion"`
	SabaticoId    int    `json:"SabaticoId"`
}
