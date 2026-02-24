package models

type FormularioSolicitud struct {
	Id          int    `orm:"column(id);pk;auto" json:"id,omitempty"`
	Contenido   string `orm:"column(contenido)" json:"contenido"`
	Activo      bool   `orm:"column(activo)" json:"activo"`
	SolicitudId int    `orm:"column(solicitud_id);rel(fk)" json:"solicitud_id"`
}
