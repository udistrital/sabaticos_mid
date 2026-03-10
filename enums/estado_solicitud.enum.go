package enums

import "strings"

// EstadoSolicitud representa los estados de una solicitud.
type EstadoSolicitud string

const (
	// BORRADOR corresponde a una solicitud en estado borrador.
	BORRADOR EstadoSolicitud = "S0"
	// RADICADA_ENVIADA_SA corresponde a una solicitud radicada y enviada a SA.
	RADICADA_ENVIADA_SA_RADICADA EstadoSolicitud = "S1"
	// RECEPCIONADA_SA corresponde a una solicitud recepcionada por SA.
	RECEPCIONADA_SA EstadoSolicitud = "S2"
	// VERIFICACION_SA corresponde a una solicitud en verificación por SA.
	VERIFICACION_SA EstadoSolicitud = "S3"
	// SUBSANACION_SOLICITADA corresponde a una solicitud con subsanación solicitada.
	SUBSANACION_SOLICITADA EstadoSolicitud = "S4"
	// TRAMITE_CONSEJO_FACULTAD corresponde a una solicitud en trámite externo Consejo de Facultad.
	TRAMITE_CONSEJO_FACULTAD EstadoSolicitud = "S5"
	// RESPUESTA_CF_REGISTRADA corresponde a una solicitud con respuesta de CF registrada.
	RESPUESTA_CF_REGISTRADA EstadoSolicitud = "S6"
	// ENVIADA_SG corresponde a una solicitud enviada a Secretaría General.
	ENVIADA_SG EstadoSolicitud = "S7"
	// RECEPCIONADA_SG corresponde a una solicitud recepcionada en Secretaría General.
	RECEPCIONADA_SG EstadoSolicitud = "S8"
	// TRAMITE_CONSEJO_ACADEMICO corresponde a una solicitud en trámite externo Consejo Académico.
	TRAMITE_CONSEJO_ACADEMICO EstadoSolicitud = "S9"
	// DECISION_CA_REGISTRADA corresponde a una solicitud con decisión de Consejo Académico registrada.
	DECISION_CA_REGISTRADA EstadoSolicitud = "S10"
	// FINALIZADA_NO_APROBADA corresponde a una solicitud finalizada no aprobada.
	FINALIZADA_NO_APROBADA EstadoSolicitud = "S11A"
	// APROBADA_PENDIENTE_RESOLUCION corresponde a una solicitud aprobada pendiente de resolución.
	APROBADA_PENDIENTE_RESOLUCION EstadoSolicitud = "S11B"
	// FINALIZADA_APROBADA_RESOLUCION corresponde a una solicitud finalizada aprobada con resolución.
	FINALIZADA_APROBADA_RESOLUCION EstadoSolicitud = "S12"
)

func ObtenerCodigoEstadoSolicitud(nombre string) (string, bool) {
	name := strings.ToUpper(strings.TrimSpace(nombre))
	switch name {
	case "BORRADOR", string(BORRADOR):
		return string(BORRADOR), true
	case "RADICADA_ENVIADA_SA", string(RADICADA_ENVIADA_SA_RADICADA):
		return string(RADICADA_ENVIADA_SA_RADICADA), true
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
