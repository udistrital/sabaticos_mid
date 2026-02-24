package models

// SolicitudCompletaResponse DTO para retornar la solicitud completa con historial y formulario
type SolicitudCompletaResponse struct {
	Solicitud  interface{} `json:"solicitud"`
	Historial  interface{} `json:"historial"`
	Formulario interface{} `json:"formulario"`
}
