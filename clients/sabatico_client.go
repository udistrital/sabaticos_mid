package clients

import (
	"fmt"

	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ConsultarSabatico(sabaticoId int) (*models.Sabatico, error) {
	var sabatico models.Sabatico
	var sabaticoRes interface{}

	if err := request.GetJson(beego.AppConfig.String("sabaticosService")+"/sabatico/"+fmt.Sprintf("%d", sabaticoId), &sabaticoRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(sabaticoRes, &sabatico); err != nil {
		return nil, err
	}

	if sabatico.Id == 0 {
		return nil, fmt.Errorf("sabatico not found: %d", sabaticoId)
	}

	return &sabatico, nil
}
