package models

// SolicitudRequest DTO para recibir la solicitud
type SolicitudRequest struct {
	TerceroId       int         `json:"TerceroId"`
	TipoSolicitudId interface{} `json:"TipoSolicitudId,omitempty"`
	SabaticoId      interface{} `json:"SabaticoId,omitempty"`
}

// SolicitudResponse DTO para retornar la solicitud
type SolicitudResponse struct {
	Id                int         `json:"Id"`
	TerceroId         int         `json:"TerceroId"`
	Activo            bool        `json:"Activo"`
	FechaCreacion     string      `json:"FechaCreacion"`
	FechaModificacion string      `json:"FechaModificacion"`
	TipoSolicitudId   interface{} `json:"TipoSolicitudId"`
	SabaticoId        interface{} `json:"SabaticoId"`
}

// Solicitud modelo de la base de datos
type Solicitud struct {
	Id                int         `orm:"column(id);pk;auto" json:"Id"`
	TerceroId         int         `orm:"column(tercero_id)" json:"TerceroId"`
	Activo            bool        `orm:"column(activo)" json:"Activo"`
	FechaCreacion     string      `orm:"column(fecha_creacion);type(timestamp without time zone)" json:"FechaCreacion"`
	FechaModificacion string      `orm:"column(fecha_modificacion);type(timestamp without time zone)" json:"FechaModificacion"`
	TipoSolicitud     interface{} `orm:"column(tipo_solicitud_id);rel(fk)" json:"TipoSolicitudId"`
	SabaticoId        interface{} `orm:"column(sabatico_id);rel(fk)" json:"SabaticoId"`
}

// SolicitudSabatico alias para compatibilidad
type SolicitudSabatico = SolicitudRequest
