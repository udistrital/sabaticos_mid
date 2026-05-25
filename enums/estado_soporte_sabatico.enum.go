package enums

import "strings"

// EstadoSabatico representa el código de abreviación del estado de un sabatico.
type EstadoSoporteSabatico string

// Códigos que vienen de tu BD (S0, S1, S2, etc.)
const (
	// BORRADOR
	PENDIENTE_REVISION_SOPORTE EstadoSoporteSabatico = "S0"

	//REVISION
	RECIBIDO_SA EstadoSoporteSabatico = "S1"
)

func ObtenerCodigoEstadoSoporteSabatico(nombre string) (string, bool) {
	name := strings.TrimSpace(nombre)
	if name == "" {
		return "", false
	}

	name = strings.ToUpper(name)
	switch name {
	case "PENDIENTE_REVISION_SOPORTE", string(PENDIENTE_REVISION_SOPORTE):
		return string(PENDIENTE_REVISION_SOPORTE), true
	case "REVISION", string(RECIBIDO_SA):
		return string(RECIBIDO_SA), true
	default:
		return "", false
	}
}
