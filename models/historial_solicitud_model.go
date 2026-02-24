package models

type HistorialSolicitud struct {
	Id                int    `orm:"column(id);pk;auto" json:"id,omitempty"`
	TerceroId         int    `orm:"column(tercero_id)" json:"tercero_id"`
	Justificacion     string `orm:"column(justificacion);null" json:"justificacion"`
	Activo            bool   `orm:"column(activo)" json:"activo"`
	EstadoSolicitudId int    `orm:"column(estado_solicitud_id);rel(fk)" json:"estado_solicitud_id"`
	SolicitudId       int    `orm:"column(solicitud_id);rel(fk)" json:"solicitud_id"`
}
