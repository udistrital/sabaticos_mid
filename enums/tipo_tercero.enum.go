package enums

import "strings"

type TipoTercero string

const (
	DOCENTE TipoTercero = "DOCENTE"
)

func ObtenerCodigoTipoTercero(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "DOCENTE":
		return string(DOCENTE), true
	default:
		return "", false
	}
}
