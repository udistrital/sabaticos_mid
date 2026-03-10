package models

// FormularioSolicitudCreateRequest estructura para crear un formulario de solicitud
type FormularioSolicitudCreateRequest struct {
	Contenido     string      `json:"Contenido"`
	SolicitudId   IdReference `json:"SolicitudId"`
	FechaCreacion string      `json:"FechaCreacion,omitempty"`
	Activo        bool        `json:"Activo"`
}
