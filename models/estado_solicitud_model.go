package models

type EstadoSolicitud struct {
	Id                int    `orm:"column(id);pk;auto"`
	CodigoAbreviacion string `orm:"column(codigo_abreviacion)"`
	Estado            string `orm:"column(estado)"`
	Descripcion       string `orm:"column(descripcion);null"`
	Activo            bool   `orm:"column(activo)"`
	FechaCreacion     string `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}
