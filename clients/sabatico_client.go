package clients

import (
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"
	"fmt"

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
		return nil, fmt.Errorf("Sábatico no Encontrado: %d", sabaticoId)
	}

	return &sabatico, nil
}
