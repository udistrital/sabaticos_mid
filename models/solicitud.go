package models

type SolicitudSabatico struct {
	Id                int
	TerceroId         int
	Activo            bool
	TipoSolicitudId   interface{}
	EstadoSolicitudId interface{}
	SabaticoId        interface{}
}
