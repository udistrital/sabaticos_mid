package enums

import "strings"

type EstadoSoporteSolicitud string

const (
	PENDIENTE  EstadoSoporteSolicitud = "PEN"
	RADICADO   EstadoSoporteSolicitud = "RAD"
	NOAPROBADO EstadoSoporteSolicitud = "NOAPROB"
	APROBADO   EstadoSoporteSolicitud = "APROB"
)

func ObtenerCodigoEstadoSoporteSolicitud(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "PENDIENTE", string(PENDIENTE):
		return string(PENDIENTE), true
	case "RADICADO", string(RADICADO):
		return string(RADICADO), true
	case "NOAPROBADO", string(NOAPROBADO):
		return string(NOAPROBADO), true
	case "APROBADO", string(APROBADO):
		return string(APROBADO), true
	default:
		return "", false
	}
}
