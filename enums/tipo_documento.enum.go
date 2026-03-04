package enums

import "strings"

// TipoDocumento representa los tipos de documentos.
type TipoDocumento string

const (
	FORMACION_ACADEMICA TipoDocumento = "FA"
)

func ObtenerCodigoTipoDocumento(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "FORMACION_ACADEMICA", string(FORMACION_ACADEMICA):
		return string(FORMACION_ACADEMICA), true
	default:
		return "", false
	}
}
