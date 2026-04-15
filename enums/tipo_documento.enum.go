package enums

import "strings"

// TipoDocumento representa los tipos de documentos.
type TipoDocumento string

const (
	SOLCITUD_SABATICO TipoDocumento = "SOL_SAB"
)

func ObtenerCodigoTipoDocumento(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "SOLICITUD_SABATICO", string(SOLCITUD_SABATICO):
		return string(SOLCITUD_SABATICO), true
	default:
		return "", false
	}
}
