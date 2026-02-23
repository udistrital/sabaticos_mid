package enums

// TipoSolicitud representa los tipos de solicitud disponibles.
type TipoSolicitud string

const (
	// NUEVA corresponde a una solicitud nueva.
	NUEVA TipoSolicitud = "NS"
	// SUSPENSION corresponde a una solicitud de suspensión.
	SUSPENSION TipoSolicitud = "SUS"
)
