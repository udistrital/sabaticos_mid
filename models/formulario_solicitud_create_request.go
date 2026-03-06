package models

// FormularioSolicitudCreateRequest estructura para crear un formulario de solicitud
type FormularioSolicitudCreateRequest struct {
	Contenido   string      `json:"Contenido"`
	SolicitudId IdReference `json:"SolicitudId"`
	Activo      bool        `json:"Activo"`
}
