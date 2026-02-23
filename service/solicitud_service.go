package service

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/time_bogota"
)

func StoreSolicitud(solicitudReq models.SolicitudRequest) (interface{}, error) {
	var respuesta interface{}
	var tercero interface{}

	errtercero := request.GetJson(beego.AppConfig.String("terceroService")+"tercero/"+fmt.Sprintf("%d", solicitudReq.TerceroId), &tercero)

	if errtercero != nil {
		return nil, errtercero
	}

	// Crear modelo con datos del request
	solicitud := models.Solicitud{
		TerceroId:         solicitudReq.TerceroId,
		TipoSolicitud:     enums.NUEVA,             // Usar código de abreviación
		SabaticoId:        solicitudReq.SabaticoId, // Puede ser nil
		Activo:            true,
		FechaCreacion:     time_bogota.TiempoBogotaFormato(),
		FechaModificacion: time_bogota.TiempoBogotaFormato(),
	}

	beego.Info("Código de tipo solicitud:", enums.NUEVA)

	errSolicitud := request.SendJson(beego.AppConfig.String("sabaticosService")+"/solicitud/", "POST", &respuesta, solicitud)

	if errSolicitud != nil {
		beego.Error("Error POST solicitud:", errSolicitud)
		return nil, errSolicitud
	}

	beego.Info("Respuesta POST solicitud:", respuesta)

	return respuesta, nil
}
