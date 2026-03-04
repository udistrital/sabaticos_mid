package enums

import "strings"

type EstadoSoporteSolicitud string

const (
	PENDIENTE EstadoSoporteSolicitud = "PEN"
)

func ObtenerCodigoEstadoSoporteSolicitud(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "PENDIENTE", string(PENDIENTE):
		return string(PENDIENTE), true
	default:
		return "", false
	}
}
