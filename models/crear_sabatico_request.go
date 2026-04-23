package models

type CrearSabaticoRequest struct {
	TerceroId      int    `json:"tercero_id"`
	Observaciones  string `json:"observaciones"`
	FechaInicio    string `json:"fecha_inicio"`
	FechaFin       string `json:"fecha_fin"`
	EstadoSabatico string `json:"estado_sabatico"`
}
