package models

// Solicitud modelo de la base de datos
type Solicitud struct {
	Id                int     `orm:"column(id);pk;auto" json:"id,omitempty"`
	TerceroId         int     `orm:"column(tercero_id)" json:"tercero_id"`
	Activo            bool    `orm:"column(activo)" json:"activo"`
	FechaCreacion     *string `orm:"column(fecha_creacion);type(timestamp without time zone)" json:"fecha_creacion,omitempty"`
	FechaModificacion *string `orm:"column(fecha_modificacion);type(timestamp without time zone)" json:"fecha_modificacion,omitempty"`
	TipoSolicitudId   int     `orm:"column(tipo_solicitud_id);rel(fk)" json:"tipo_solicitud_id"`
	SabaticoId        *int    `orm:"column(sabatico_id);rel(fk)" json:"sabatico_id,omitempty"`
}
