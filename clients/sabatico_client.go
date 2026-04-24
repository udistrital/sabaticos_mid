package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/udistrital/sabaticos_mid/enums"
	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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

func RegistrarSabatico(solicitudId int, terceroId int, observaciones string, fechaInicio string, fechaFin string, estadoSabatico string) (*models.CrearSabaticoResult, error) {
	crudURL := strings.TrimRight(beego.AppConfig.String("sabaticosService"), "/") + "/sabatico"

	soportes, err := ConsultarSoportesSolicitud(solicitudId)

	if err != nil && len(soportes) > 0 {
		return nil, fmt.Errorf("error consultando soportes de la solicitud %d: %v", solicitudId, err)
	}

	estadoSabaticoId, err := ConsultarIdEstadoSabatico(estadoSabatico)
	if err != nil {
		return nil, fmt.Errorf("error consultando id de estado sabático: %v", err)
	}

	payload := map[string]interface{}{
		"Activo": true,
		"EstadoSabaticoId": map[string]interface{}{
			"Id": estadoSabaticoId,
		},
		"FechaCreacion":     time.Now().Format("2006-01-02 15:04:05"),
		"FechaFin":          fechaFin,
		"FechaInicio":       fechaInicio,
		"FechaModificacion": time.Now().Format("2006-01-02 15:04:05"),
		"Observaciones":     observaciones,
		"TerceroId":         terceroId,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error serializando payload de sabático: %v", err)
	}

	req, err := http.NewRequest(
		"POST",
		crudURL,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error creando request al CRUD: %v",
			err,
		)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		logs.Error(
			"error consumiendo sabaticos_crud:",
			err,
		)

		return nil, fmt.Errorf(
			"error consumiendo sabaticos_crud: %v",
			err,
		)
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"error leyendo respuesta del CRUD: %v",
			err,
		)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf(
			"sabaticos_crud respondió con estado %d: %s",
			resp.StatusCode,
			string(respBytes),
		)
	}

	/*
		Extraer correctamente Data del response OAS
	*/
	var response interface{}
	var result models.CrearSabaticoResult

	if err := json.Unmarshal(respBytes, &response); err != nil {
		return nil, fmt.Errorf(
			"error decodificando respuesta del CRUD: %v",
			err,
		)
	}

	if err := helpers.ExtractDataApi(response, &result); err != nil {
		return nil, fmt.Errorf(
			"error extrayendo data de respuesta CRUD: %v",
			err,
		)
	}

	fmt.Println("Sabatico creado:", result)

	return &result, nil
}

/*
   Helper functions
*/

/*
ConsultarIdEstadoSabatico is intended to obtain the ID of a sabbatical status
based on its abbreviation code or the status name.
*/
func ConsultarIdEstadoSabatico(estado string) (int, error) {
	var estadoSabaticoRes interface{}
	var estados []models.EstadoSabatico

	codigo, ok := enums.ObtenerCodigoEstadoSabatico(estado)
	if !ok {
		codigo = strings.TrimSpace(estado)
	}

	baseURL := strings.TrimRight(beego.AppConfig.String("sabaticosService"), "/")
	if baseURL == "" {
		return 0, fmt.Errorf("la configuración 'sabaticosService' no está definida")
	}

	url := baseURL + "/estado_sabatico?query=Activo:true,CodigoAbreviacion:" + codigo + "&limit=1"

	if err := request.GetJson(url, &estadoSabaticoRes); err != nil {
		return 0, err
	}

	if err := helpers.ExtractDataApi(estadoSabaticoRes, &estados); err != nil {
		return 0, err
	}

	if len(estados) == 0 {
		url = baseURL + "/estado_sabatico?query=Activo:true,NombreEstado:" + strings.TrimSpace(estado) + "&limit=1"

		if err := request.GetJson(url, &estadoSabaticoRes); err != nil {
			return 0, err
		}

		if err := helpers.ExtractDataApi(estadoSabaticoRes, &estados); err != nil {
			return 0, err
		}
	}

	if len(estados) == 0 {
		return 0, fmt.Errorf("no se encontró estado_sabatico para valor '%s'", estado)
	}

	return estados[0].Id, nil
}
