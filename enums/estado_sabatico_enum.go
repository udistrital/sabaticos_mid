package enums

import "strings"

// EstadoSabatico representa el código de abreviación del estado de un sabático.
type EstadoSabatico string

const (
	// Estado base durante el desarrollo del plan de trabajo
	EN_EJECUCION EstadoSabatico = "ES0"

	// El sabático entra en trámite de modificación del plan de trabajo
	CARGUE_PLAN_TRABAJO EstadoSabatico = "ES1"

	// El sabático entra en revision de los secretarios
	REVISION_SA EstadoSabatico = "ES2"

	//Socializacion Pendiente
	SOCIALIZACION_PENDIENTE EstadoSabatico = "ES3"

	//SUBSANACION
	SUBSANACION EstadoSabatico = "ES4"
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
	case "CARGUE_PLAN_TRABAJO", string(CARGUE_PLAN_TRABAJO):
		return string(CARGUE_PLAN_TRABAJO), true
	case "REVISION_SA", string(REVISION_SA):
		return string(REVISION_SA), true
	case "SOCIALIZACION_PENDIENTE", string(SOCIALIZACION_PENDIENTE):
		return string(SOCIALIZACION_PENDIENTE), true
	case "SUBSANACION", string(SUBSANACION):
		return string(SUBSANACION), true
	default:
		return "", false
	}
}
