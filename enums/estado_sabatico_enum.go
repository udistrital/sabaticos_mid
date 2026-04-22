package enums

import "strings"

// EstadoSabatico representa el código de abreviación del estado de un sabático.
type EstadoSabatico string

const (
	// Estado base durante el desarrollo del plan de trabajo
	EN_EJECUCION EstadoSabatico = "ES0"

	// El sabático entra en trámite de modificación del plan de trabajo
	A_MODIFICAR EstadoSabatico = "ES1"

	// El producto académico está siendo evaluado/revisado
	PRODUCTO_EN_REVISION EstadoSabatico = "ES2"

	// El producto ya fue aprobado para pasar a socialización
	PENDIENTE_SOCIALIZACION EstadoSabatico = "ES3"

	// El producto ya fue socializado
	SOCIALIZADO EstadoSabatico = "ES4"

	// El sabático queda suspendido
	SUSPENDIDO EstadoSabatico = "ES5"

	// El proceso de seguimiento del sabático terminó
	FINALIZADO EstadoSabatico = "ES6"
)

func ObtenerCodigoEstadoSabatico(nombre string) (string, bool) {
	name := strings.TrimSpace(nombre)
	if name == "" {
		return "", false
	}

	name = strings.ToUpper(name)

	switch name {
	case "EN_EJECUCION", string(EN_EJECUCION):
		return string(EN_EJECUCION), true
	case "A_MODIFICAR", string(A_MODIFICAR):
		return string(A_MODIFICAR), true
	case "PRODUCTO_EN_REVISION", string(PRODUCTO_EN_REVISION):
		return string(PRODUCTO_EN_REVISION), true
	case "PENDIENTE_SOCIALIZACION", string(PENDIENTE_SOCIALIZACION):
		return string(PENDIENTE_SOCIALIZACION), true
	case "SOCIALIZADO", string(SOCIALIZADO):
		return string(SOCIALIZADO), true
	case "SUSPENDIDO", string(SUSPENDIDO):
		return string(SUSPENDIDO), true
	case "FINALIZADO", string(FINALIZADO):
		return string(FINALIZADO), true
	default:
		return "", false
	}
}
