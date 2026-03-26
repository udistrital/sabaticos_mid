package enums

import "strings"

// EstadoSolicitud representa el código de abreviación del estado de una solicitud.
type EstadoSolicitud string

// Códigos que vienen de tu BD (S0, S1, S2, etc.)
const (
	// BORRADOR
	BORRADOR EstadoSolicitud = "S0"

	// Estado inicial una vez el docente radica y envía a SA
	RADICADA_ENVIADA_SA EstadoSolicitud = "S1"

	// SA recibe formalmente la solicitud
	RECEPCIONADA_SA EstadoSolicitud = "S2"

	// SA está verificando el checklist y documentos
	VERIFICACION_SA EstadoSolicitud = "S3"

	// SA pide subsanaciones
	SUBSANACION_SOLICITADA EstadoSolicitud = "S4"

	// Solicitud en trámite externo en Consejo de Facultad
	TRAMITE_CONSEJO_FACULTAD EstadoSolicitud = "S5"

	// SA registra el acta o concepto del CF
	RESPUESTA_CF_REGISTRADA EstadoSolicitud = "S6"

	// Se envía expediente a Secretaría General
	ENVIADA_SG EstadoSolicitud = "S7"

	// SG recibe formalmente
	RECEPCIONADA_SG EstadoSolicitud = "S8"

	// Solicitud en trámite externo en Consejo Académico
	TRAMITE_CONSEJO_ACADEMICO EstadoSolicitud = "S9"

	// SG registra la decisión de Consejo Académico
	DECISION_CA_REGISTRADA EstadoSolicitud = "S10"

	// Solicitud cerrada por decisión negativa
	FINALIZADA_NO_APROBADA EstadoSolicitud = "S11A"

	// Solicitud aprobada, pendiente de resolución emitida
	APROBADA_PENDIENTE_RESOLUCION EstadoSolicitud = "S11B"

	// Solicitud cerrada con resolución emitida
	FINALIZADA_APROBADA_RESOLUCION EstadoSolicitud = "S12"
)

// Uso de switch, igual a tu estructura de `EstadoSoporteSolicitud`
func ObtenerCodigoEstadoSolicitud(nombre string) (string, bool) {
	name := strings.TrimSpace(nombre)
	if name == "" {
		return "", false
	}

	name = strings.ToUpper(name)

	switch name {
	case "BORRADOR", string(BORRADOR):
		return string(BORRADOR), true
	case "RADICADA_ENVIADA_SA", string(RADICADA_ENVIADA_SA):
		return string(RADICADA_ENVIADA_SA), true
	case "RECEPCIONADA_SA", string(RECEPCIONADA_SA):
		return string(RECEPCIONADA_SA), true
	case "VERIFICACION_SA", string(VERIFICACION_SA):
		return string(VERIFICACION_SA), true
	case "SUBSANACION_SOLICITADA", string(SUBSANACION_SOLICITADA):
		return string(SUBSANACION_SOLICITADA), true
	case "TRAMITE_CONSEJO_FACULTAD", string(TRAMITE_CONSEJO_FACULTAD):
		return string(TRAMITE_CONSEJO_FACULTAD), true
	case "RESPUESTA_CF_REGISTRADA", string(RESPUESTA_CF_REGISTRADA):
		return string(RESPUESTA_CF_REGISTRADA), true
	case "ENVIADA_SG", string(ENVIADA_SG):
		return string(ENVIADA_SG), true
	case "RECEPCIONADA_SG", string(RECEPCIONADA_SG):
		return string(RECEPCIONADA_SG), true
	case "TRAMITE_CONSEJO_ACADEMICO", string(TRAMITE_CONSEJO_ACADEMICO):
		return string(TRAMITE_CONSEJO_ACADEMICO), true
	case "DECISION_CA_REGISTRADA", string(DECISION_CA_REGISTRADA):
		return string(DECISION_CA_REGISTRADA), true
	case "FINALIZADA_NO_APROBADA", string(FINALIZADA_NO_APROBADA):
		return string(FINALIZADA_NO_APROBADA), true
	case "APROBADA_PENDIENTE_RESOLUCION", string(APROBADA_PENDIENTE_RESOLUCION):
		return string(APROBADA_PENDIENTE_RESOLUCION), true
	case "FINALIZADA_APROBADA_RESOLUCION", string(FINALIZADA_APROBADA_RESOLUCION):
		return string(FINALIZADA_APROBADA_RESOLUCION), true
	default:
		return "", false
	}
}
