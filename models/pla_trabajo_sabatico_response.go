package models

import "time"

type PlanTrabajoSabaticoResponse struct {
	SolicitudId   int       `json:"SolicitudId"`
	TerceroId     int       `json:"TerceroId"`
	Observaciones string    `json:"Observaciones"`
	FechaInicio   time.Time `json:"FechaInicio"`
	FechaFin      time.Time `json:"FechaFin"`
}
