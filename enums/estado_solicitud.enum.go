package enums

import "strings"

// EstadoSolicitud representa los estados de una solicitud.
type EstadoSolicitud string

const (
	// ENVIADA corresponde a una solicitud enviada.
	ENVIADA EstadoSolicitud = "ENV"
	// APROBADA corresponde a una solicitud aprobada.
	APROBADA EstadoSolicitud = "APR"
)

func ObtenerCodigoEstadoSolicitud(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "ENVIADA", string(ENVIADA):
		return string(ENVIADA), true
	case "APROBADA", string(APROBADA):
		return string(APROBADA), true
	default:
		return "", false
	}
}
