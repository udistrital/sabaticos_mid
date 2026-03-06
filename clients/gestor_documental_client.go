package clients

import (
	"fmt"
	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarGestorDocumental(idTipoDocumento int, nombre string, descripcion string, metadatos interface{}, file string) (*models.GestorDocumental, error) {
	var gestorDocumentalRes interface{}
	var gestorDocumentalFinal models.GestorDocumental
	var gestorDocumentalFinalList []models.GestorDocumental

	gestorDocumental := models.GestorDocumentalCreateRequest{
		IdTipoDocumento: idTipoDocumento,
		Nombre:          nombre,
		Descripcion:     descripcion,
		Metadatos:       metadatos,
		File:            file,
	}

	payload := []models.GestorDocumentalCreateRequest{gestorDocumental}

	if err := request.SendJson(beego.AppConfig.String("gestorDocumentalService")+"/document/uploadAnyFormat", "POST", &gestorDocumentalRes, payload); err != nil {
		return nil, fmt.Errorf("falló uploadAnyFormat en gestor documental: %w", err)
	}

	if err := helpers.ValidateServiceResponse(gestorDocumentalRes); err != nil {
		return nil, fmt.Errorf("gestor documental devolvió error: %w", err)
	}

	if err := helpers.ExtractDataApi(gestorDocumentalRes, &gestorDocumentalFinalList); err == nil && len(gestorDocumentalFinalList) > 0 {
		return &gestorDocumentalFinalList[0], nil
	}

	if err := helpers.ExtractDataApi(gestorDocumentalRes, &gestorDocumentalFinal); err != nil {
		return nil, err
	}

	return &gestorDocumentalFinal, nil

}

func ConsultarTipoDocumento(codigoAbreviacion string) (*models.TipoDocumento, error) {
	var tipoDocumentoRes interface{}
	var tipoDocumento []models.TipoDocumento

	if err := request.GetJson(beego.AppConfig.String("documentosService")+"tipo_documento?query=CodigoAbreviacion:"+codigoAbreviacion, &tipoDocumentoRes); err != nil {
		return nil, err
	}
	if err := helpers.ExtractDataApi(tipoDocumentoRes, &tipoDocumento); err != nil {
		return nil, err
	}

	if len(tipoDocumento) == 0 {
		return nil, fmt.Errorf("Tipo de Documento no Encontrado: %s", codigoAbreviacion)
	}

	return &tipoDocumento[0], nil

}
