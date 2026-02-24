package clients

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarSolicitud(solicitudReq models.SolicitudRequest) (*models.Solicitud, error) {
	var solicitudRes interface{}

	tipoSolicitudId := helpers.InterfaceToInt(solicitudReq.TipoSolicitudId, int(enums.NUEVA))
	sabaticoId := helpers.InterfaceToIntPtr(solicitudReq.SabaticoId)

	solicitud := models.Solicitud{
		TerceroId:       solicitudReq.TerceroId,
		TipoSolicitudId: tipoSolicitudId,
		SabaticoId:      sabaticoId,
		Activo:          true,
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/solicitud/", "POST", &solicitudRes, solicitud); err != nil {
		beego.Error("Error POST solicitud:", err)
		return nil, err
	}

	var solicitudCreada *models.Solicitud
	if err := helpers.ExtractDataApi(solicitudRes, &solicitudCreada); err != nil {
		beego.Error("Error extrayendo datos de solicitud:", err)
		return nil, err
	}

	beego.Info("ID de solicitud creada:", solicitudCreada.Id)
	return solicitudCreada, nil
}
