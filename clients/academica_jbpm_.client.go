package clients

import (
	"errors"

	"github.com/astaxie/beego"

	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/utils_oas/request"
)

func ConsultarSecretariaAcademicaDocumentoUserId(documentoId string) (*models.Persona, error) {
	var secretariaResp models.SecretariaAcademicaResponse

	if err := request.GetXml(beego.AppConfig.String("academicaJbpmService")+"secretaria_academica/"+documentoId, &secretariaResp); err != nil {
		return nil, err
	}

	if secretariaResp.Persona == nil {
		return nil, errors.New("no persona found in response")
	}

	return secretariaResp.Persona, nil
}
