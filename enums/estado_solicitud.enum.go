package enums

import "strings"

// EstadoSolicitud representa los estados de una solicitud.
type EstadoSolicitud string

const (
	// ENVIADA corresponde a una solicitud enviada.
	ENVIADA EstadoSolicitud = "ENV"
	// RADICADA corresponde a una solicitud radicada.
	RADICADA EstadoSolicitud = "RAD"
)

func ObtenerCodigoEstadoSolicitud(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "ENVIADA", string(ENVIADA):
		return string(ENVIADA), true
	case "RADICADA", string(RADICADA):
		return string(RADICADA), true
	default:
		return "", false
	}
}
