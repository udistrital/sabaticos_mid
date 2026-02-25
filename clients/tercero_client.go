package clients

import (
	"api_mid_sabaticos/enums"
	"api_mid_sabaticos/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ValidarTercero(terceroId int) error {
	tipoTercero, err := ConsultarTipoTercero(enums.DOCENTE)
	if err != nil {
		return err
	}

	terceroTipoTercero, err := ConsultaTerceroTipoTercero(terceroId, tipoTercero.Id)
	if err != nil {
		return err
	}

	if terceroTipoTercero == nil {
		return fmt.Errorf("no se encontro ese tercero con rol de docente")
	}

	fmt.Printf("%+v\n", terceroTipoTercero)

	return nil
}

func ConsultarTipoTercero(codigoAbreviacion enums.TipoTercero) (*models.TipoTercero, error) {
	var tipos []models.TipoTercero

	if err := request.GetJson(beego.AppConfig.String("terceroService")+"tipo_tercero?query="+fmt.Sprintf("CodigoAbreviacion:%s", string(codigoAbreviacion)), &tipos); err != nil {
		beego.Error("Error GET tipo_tercero:", err)
		return nil, err
	}

	if len(tipos) == 0 {
		return nil, fmt.Errorf("tipo tercero no encontrado: %s", string(codigoAbreviacion))
	}

	return &tipos[0], nil
}

func ConsultaTerceroTipoTercero(terceroId int, tipoTerceroId int) (*models.TerceroTipoTercero, error) {
	var tercerosTipoTerceros []models.TerceroTipoTercero

	if err := request.GetJson(beego.AppConfig.String("terceroService")+"tercero_tipo_tercero?query="+fmt.Sprintf("TerceroId.Id:%d,TipoTerceroId.Id:%d", terceroId, tipoTerceroId), &tercerosTipoTerceros); err != nil {
		beego.Error("Error GET tercero_tipo_tercero:", err)
		return nil, err
	}

	if len(tercerosTipoTerceros) == 0 {
		return nil, nil
	}

	first := tercerosTipoTerceros[0]
	if first.Id == 0 && first.TerceroId == nil && first.TipoTerceroId == nil {
		return nil, nil
	}

	return &first, nil
}
