package enums

import "strings"

type TipoSolicitud string

const (
	NUEVA      TipoSolicitud = "NS"
	SUSPENSION TipoSolicitud = "SS"
)

func ObtenerCodigoTipoSolicitud(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "NUEVA", string(NUEVA):
		return string(NUEVA), true
	case "SUSPENSION", string(SUSPENSION):
		return string(SUSPENSION), true
	default:
		return "", false
	}
}
