package clients

import (
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarSolicitud(terceroId int, codigoTipoSolicitud string, sabaticoId *int) (*models.Solicitud, error) {
	var solicitudRes interface{}
	var solicitudCreada *models.Solicitud

	tipoSolicitud, err := ConsultarTipoSolicitud(codigoTipoSolicitud)
	if err != nil {
		return nil, err
	}

	solicitud := models.SolicitudCreateRequest{
		TerceroId: terceroId,
		Activo:    true,
		TipoSolicitudId: models.IdReference{
			Id: tipoSolicitud.Id,
		},
	}

	if err := request.SendJson(beego.AppConfig.String("sabaticosService")+"/solicitud/", "POST", &solicitudRes, solicitud); err != nil {
		beego.Error("Error POST solicitud:", err)
		return nil, err
	}

	if err := helpers.ExtractDataApi(solicitudRes, &solicitudCreada); err != nil {
		beego.Error("Error extrayendo datos de solicitud:", err)
		return nil, err
	}

	beego.Info("ID de solicitud creada:", solicitudCreada.Id)
	return solicitudCreada, nil
}
