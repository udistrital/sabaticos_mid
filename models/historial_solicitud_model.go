package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type HistorialSolicitud struct {
	Id                int        `orm:"column(id);pk;auto" json:"id,omitempty"`
	TerceroId         int        `orm:"column(tercero_id)" json:"tercero_id"`
	Justificacion     string     `orm:"column(justificacion);null" json:"justificacion"`
	Activo            bool       `orm:"column(activo)" json:"activo"`
	EstadoSolicitudId int        `orm:"column(estado_solicitud_id);rel(fk)" json:"estado_solicitud_id"`
	SolicitudId       int        `orm:"column(solicitud_id);rel(fk)" json:"solicitud_id"`
	FechaModificacion *time.Time `orm:"column(fecha_modificacion);type(datetime);null" json:"fecha_modificacion,omitempty"`
}

func GetHistorialSolicitudActual(solicitudId int) (*HistorialSolicitud, error) {
	if solicitudId <= 0 {
		return nil, errors.New("solicitud_id inválido")
	}

	o := orm.NewOrm()

	var h HistorialSolicitud
	err := o.QueryTable(new(HistorialSolicitud)).
		Filter("SolicitudId", solicitudId).
		OrderBy("-FechaModificacion", "-Id").
		Limit(1).
		One(&h)

	if err == orm.ErrNoRows {
		return nil, orm.ErrNoRows
	}
	if err != nil {
		return nil, err
	}
	return &h, nil
}
