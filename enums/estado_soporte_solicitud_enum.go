package enums

import "strings"

// EstadoSoporteSolicitud representa el código de abreviación de soporte.
type EstadoSoporteSolicitud string

const (
	PENDIENTE_REVISION       EstadoSoporteSolicitud = "PEN"    // Pendiente de revisión radicado
	SUBSANADO_DOCENTE        EstadoSoporteSolicitud = "SUBS"   // Subsanado por docente reenviado
	SA_RECIBIDO_SA           EstadoSoporteSolicitud = "SAREC"  // Recibido por Secretaría Académica
	SA_PENDIENTE_REVISION_SA EstadoSoporteSolicitud = "SAPEND" // Pendiente revisión SA checklist
	SA_INVALIDO              EstadoSoporteSolicitud = "SAINV"  // Inválido según SA requiere subsanación
	SA_VALIDO                EstadoSoporteSolicitud = "SAOK"   // Válido según SA
	SG_RECIBIDO_SG           EstadoSoporteSolicitud = "SGREC"  // Recibido por Secretaría General
	SG_PENDIENTE_REVISION_SG EstadoSoporteSolicitud = "SGPEND" // Pendiente revisión SG
	SG_INVALIDO              EstadoSoporteSolicitud = "SGINV"  // Inválido según SG
	SG_VALIDO                EstadoSoporteSolicitud = "SGOK"   // Válido según SG
)

func ObtenerCodigoEstadoSoporteSolicitud(nombre string) (string, bool) {
	name := strings.TrimSpace(nombre)
	if name == "" {
		return "", false
	}

	name = strings.ToUpper(name)

	switch name {
	case "PENDIENTE_REVISION", string(PENDIENTE_REVISION):
		return string(PENDIENTE_REVISION), true
	case "SUBSANADO_DOCENTE", string(SUBSANADO_DOCENTE):
		return string(SUBSANADO_DOCENTE), true
	case "SA_RECIBIDO_SA", string(SA_RECIBIDO_SA):
		return string(SA_RECIBIDO_SA), true
	case "SA_PENDIENTE_REVISION_SA", string(SA_PENDIENTE_REVISION_SA):
		return string(SA_PENDIENTE_REVISION_SA), true
	case "SA_INVALIDO", string(SA_INVALIDO):
		return string(SA_INVALIDO), true
	case "SA_VALIDO", string(SA_VALIDO):
		return string(SA_VALIDO), true
	case "SG_RECIBIDO_SG", string(SG_RECIBIDO_SG):
		return string(SG_RECIBIDO_SG), true
	case "SG_PENDIENTE_REVISION_SG", string(SG_PENDIENTE_REVISION_SG):
		return string(SG_PENDIENTE_REVISION_SG), true
	case "SG_INVALIDO", string(SG_INVALIDO):
		return string(SG_INVALIDO), true
	case "SG_VALIDO", string(SG_VALIDO):
		return string(SG_VALIDO), true
	default:
		return "", false
	}
}
