package models

// SolicitudResponse DTO para retornar la solicitud
type SolicitudResponse struct {
	Solicitud  interface{} `json:"solicitud"`
	Historial  interface{} `json:"historial"`
	Formulario interface{} `json:"formulario"`
}
