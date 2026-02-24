package enums

// EstadoSolicitud representa los estados de una solicitud.
type EstadoSolicitud int

const (
	// ENVIADA corresponde a una solicitud enviada.
	ENVIADA EstadoSolicitud = 2
	// APROBADA corresponde a una solicitud aprobada.
	APROBADA EstadoSolicitud = 1
	// RECHAZADA corresponde a una solicitud rechazada.
	RECHAZADA EstadoSolicitud = 3
	// PENDIENTE corresponde a una solicitud pendiente.
	PENDIENTE EstadoSolicitud = 4
	// CANCELADA corresponde a una solicitud cancelada.
	CANCELADA EstadoSolicitud = 5
)
