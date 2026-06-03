package models

type PlanTrabajoSabaticoRequest struct {
	Justificacion string `json:"Justificacion"`
	SabaticoId    int    `json:"SabaticoId"`
}
