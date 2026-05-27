package enums

import "strings"

type TipoSolicitud string

const (
	NUEVA        TipoSolicitud = "NS"
	SUSPENSION   TipoSolicitud = "SS"
	MODIFICACION TipoSolicitud = "MS"
)

func ObtenerCodigoTipoSolicitud(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "NUEVA", string(NUEVA):
		return string(NUEVA), true
	case "SUSPENSION", string(SUSPENSION):
		return string(SUSPENSION), true
	case "MODIFICACION", "MODIFICATION", string(MODIFICACION):
		return string(MODIFICACION), true
	default:
		return "", false
	}
}
