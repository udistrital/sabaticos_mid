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

func ConsultarEstadoSoporteSabatico(codigo string) (*models.EstadoSoporteSabatico, error) {
	var estadoSoporteSabaticoRes interface{}
	var estadoSoporteSabatico []models.EstadoSoporteSabatico

	codigoAbreviacion, ok := enums.ObtenerCodigoEstadoSoporteSabatico(codigo)
	if !ok {
		return nil, fmt.Errorf("request status not valid: %s", codigo)
	}

	baseURL := beego.AppConfig.String("sabaticosService")
	if baseURL == "" {
		return nil, fmt.Errorf("config sabaticosService is empty")
	}

	url := baseURL + "/estado_soporte_sabatico?query=Activo:true,CodigoAbreviacion:" + codigoAbreviacion

	if err := request.GetJson(url, &estadoSoporteSabaticoRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(estadoSoporteSabaticoRes, &estadoSoporteSabatico); err != nil {
		return nil, err
	}

	if len(estadoSoporteSabatico) == 0 {
		return nil, fmt.Errorf("request status not found: %s", codigo)
	}

	return &estadoSoporteSabatico[0], nil
}

func ConsultarSoportesSabaticos(sabaticoId int) ([]models.SoporteSabatico, error) {
	var SoportesSabaticoRes interface{}
	var SoportesSabatico []models.SoporteSabatico

	url := beego.AppConfig.String("sabaticosService") + "soporte_sabatico?query=Activo:true,SabaticoId:" + fmt.Sprint(sabaticoId)

	if err := request.GetJson(url, &SoportesSabaticoRes); err != nil {
		return nil, err
	}

	if err := helpers.ExtractDataApi(SoportesSabaticoRes, &SoportesSabatico); err != nil {
		return nil, err
	}

	return SoportesSabatico, nil
}

func ConsultarporEstadoHistorialEstadoSabatico(sabaticoId int, estadoSabatico enums.EstadoSabatico) ([]models.HistorialEstadoSabatico, error) {
	var response interface{}
	var historial []models.HistorialEstadoSabatico

	url := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + fmt.Sprintf("/historial_estado_sabatico?query=SabaticoId.Id:%d,Activo:true,EstadoSabaticoId.CodigoAbreviacion:%s&limit=-1", sabaticoId, string(estadoSabatico))

	if err := request.GetJson(
		url,
		&response,
	); err != nil {
		return nil, fmt.Errorf(
			"error consumiendo historial_estado_sabatico: %v",
			err,
		)
	}

	if err := helpers.ExtractDataApi(
		response,
		&historial,
	); err != nil {
		return nil, fmt.Errorf(
			"error extrayendo historial estado sabatico: %v",
			err,
		)
	}

	return historial, nil
}

func ConsultarTodosHistorialEstadoSabatico(sabaticoId int) ([]models.HistorialEstadoSabatico, error) {
	var response interface{}
	var historial []models.HistorialEstadoSabatico

	url := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + fmt.Sprintf("/historial_estado_sabatico?query=SabaticoId.Id:%d", sabaticoId)

	if err := request.GetJson(
		url,
		&response,
	); err != nil {
		return nil, fmt.Errorf(
			"error consumiendo historial_estado_sabatico: %v",
			err,
		)
	}

	if err := helpers.ExtractDataApi(
		response,
		&historial,
	); err != nil {
		return nil, fmt.Errorf(
			"error extrayendo historial estado sabatico: %v",
			err,
		)
	}

	return historial, nil
}

func ConsultarHistorialEstadoSabatico(sabaticoId int) ([]models.HistorialEstadoSabatico, error) {
	var response interface{}
	var historial []models.HistorialEstadoSabatico

	url := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + fmt.Sprintf("/historial_estado_sabatico?query=SabaticoId.Id:%d,Activo:true", sabaticoId)

	if err := request.GetJson(
		url,
		&response,
	); err != nil {
		return nil, fmt.Errorf(
			"error consumiendo historial_estado_sabatico: %v",
			err,
		)
	}

	if err := helpers.ExtractDataApi(
		response,
		&historial,
	); err != nil {
		return nil, fmt.Errorf(
			"error extrayendo historial estado sabatico: %v",
			err,
		)
	}

	return historial, nil
}

func ConsultarPlanTrabajoSabatico(HistorialEstadoSabaticoId int) ([]models.HistorialEstadoSabatico, error) {
	var response interface{}
	var historial []models.HistorialEstadoSabatico

	//COLOCAR ENUM PARA EL ESTADO SABATICO EN EL QUERY
	url := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + fmt.Sprintf("/historial_estado_sabatico?query=Id:%d,EstadoSabaticoId.CodigoAbreviacion:ES1", HistorialEstadoSabaticoId)

	if err := request.GetJson(
		url,
		&response,
	); err != nil {
		return nil, fmt.Errorf(
			"error consumiendo historial_estado_sabatico: %v",
			err,
		)
	}

	if err := helpers.ExtractDataApi(
		response,
		&historial,
	); err != nil {
		return nil, fmt.Errorf(
			"error extrayendo historial estado sabatico: %v",
			err,
		)
	}

	return historial, nil
}

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

func RegistrarSabatico(
	solicitudId int,
	terceroId int,
	observaciones string,
	fechaInicio string,
	fechaFin string,
	estadoSabatico string,
) (*models.CrearSabaticoResult, error) {

	crudURL := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + "/sabatico"

	payload := map[string]interface{}{
		"Activo": true,

		"FechaCreacion": time.Now().Format(
			"2006-01-02 15:04:05",
		),

		"FechaFin": fechaFin,

		"FechaInicio": fechaInicio,

		"FechaModificacion": time.Now().Format(
			"2006-01-02 15:04:05",
		),

		"Observaciones": observaciones,

		"TerceroId": terceroId,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf(
			"error serializando payload de sabático: %v",
			err,
		)
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

	var response interface{}
	var result models.CrearSabaticoResult

	if err := json.Unmarshal(
		respBytes,
		&response,
	); err != nil {

		return nil, fmt.Errorf(
			"error decodificando respuesta CRUD: %v",
			err,
		)
	}

	if err := helpers.ExtractDataApi(
		response,
		&result,
	); err != nil {

		return nil, fmt.Errorf(
			"error extrayendo data CRUD: %v",
			err,
		)
	}

	/*
		Crear historial estado sabático
	*/

	estadoSabaticoId, err := ConsultarIdEstadoSabatico(
		estadoSabatico,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error consultando id estado sabático: %v",
			err,
		)
	}

	_, err = CrearHistorialEstadoSabatico(
		terceroId,
		"Creación inicial del sabático",
		estadoSabaticoId,
		result.Id,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error creando historial estado sabático: %v",
			err,
		)
	}

	return &result, nil
}

func RegistrarSoporteSabatico(
	gestorDocumentalId int,
	sabaticoId int,
	estadoSoporteSabaticoId int,
	rolUsuario string,
) (*models.SoporteSabatico, error) {
	var response interface{}
	var soporteSabatico models.SoporteSabatico

	SoporteSabaticoCreateRequest := models.SoporteSabaticoCreateRequest{
		Activo:                  true,
		FechaCreacion:           time.Now().Format("2006-01-02 15:04:05"),
		FechaModificacion:       time.Now().Format("2006-01-02 15:04:05"),
		DocumentoId:             gestorDocumentalId,
		SabaticoId:              models.IdReference{Id: sabaticoId},
		EstadoSoporteSabaticoId: models.IdReference{Id: estadoSoporteSabaticoId},
		RolUsuario:              rolUsuario,
	}

	url := beego.AppConfig.String("sabaticosService") + "/soporte_sabatico"

	if err := request.SendJson(
		url,
		"POST",
		&response,
		SoporteSabaticoCreateRequest,
	); err != nil {
		return nil, fmt.Errorf(
			"error consumiendo soporte_sabatico: %v",
			err,
		)
	}
	if err := helpers.ExtractDataApi(
		response,
		&soporteSabatico,
	); err != nil {
		beego.Error("error extracting support request data:", err)
		return nil, err
	}

	return &soporteSabatico, nil
}

func CrearHistorialEstadoSabatico(
	terceroId int,
	justificacion string,
	estadoSabaticoId int,
	sabaticoId int,
) (*models.HistorialEstadoSabatico, error) {

	var response interface{}
	var historial models.HistorialEstadoSabatico

	url := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + "/historial_estado_sabatico"

	payload := map[string]interface{}{
		"TerceroId":         terceroId,
		"Justificacion":     justificacion,
		"Activo":            true,
		"FechaCreacion":     time.Now().Format("2006-01-02 15:04:05"),
		"FechaModificacion": time.Now().Format("2006-01-02 15:04:05"),
		"EstadoSabaticoId": map[string]interface{}{
			"Id": estadoSabaticoId,
		},
		"SabaticoId": map[string]interface{}{
			"Id": sabaticoId,
		},
	}

	logs.Info(
		"payload crear historial estado sabatico: %+v",
		payload,
	)

	if err := request.SendJson(
		url,
		"POST",
		&response,
		payload,
	); err != nil {

		return nil, fmt.Errorf(
			"error consumiendo historial_estado_sabatico: %v",
			err,
		)
	}

	if err := helpers.ValidateServiceResponse(response); err != nil {
		return nil, fmt.Errorf(
			"sabaticosService historial_estado_sabatico returned error: %w",
			err,
		)
	}

	if err := helpers.ExtractDataApi(
		response,
		&historial,
	); err != nil {

		return nil, fmt.Errorf(
			"error extrayendo historial estado sabatico: %v",
			err,
		)
	}

	return &historial, nil
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

	fmt.Println("-------")
	fmt.Println(url)
	fmt.Println("-------")

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

func DesactivarHistorialEstadoSabatico(historial models.HistorialEstadoSabatico) error {
	var response interface{}
	url := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + fmt.Sprintf("/historial_estado_sabatico/%d", historial.Id)

	historialEstadoSabatico := models.HistorialEstadoSabatico{
		Id:                historial.Id,
		TerceroId:         historial.TerceroId,
		Justificacion:     historial.Justificacion,
		Activo:            false,
		FechaCreacion:     historial.FechaCreacion,
		FechaModificacion: historial.FechaModificacion,
		EstadoSabaticoId:  historial.EstadoSabaticoId,
		SabaticoId:        historial.SabaticoId,
	}

	if err := request.SendJson(
		url,
		"PUT",
		&response,
		historialEstadoSabatico,
	); err != nil {
		return fmt.Errorf(
			"error desactivando historial_estado_sabatico: %v",
			err,
		)
	}

	if err := helpers.ValidateServiceResponse(response); err != nil {
		return fmt.Errorf(
			"sabaticosService historial_estado_sabatico returned error: %w",
			err,
		)
	}

	return nil
}

func ActualizarHistorialEstadoSabatico(historial models.HistorialEstadoSabatico) error {
	var response interface{}
	url := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + fmt.Sprintf("/historial_estado_sabatico/%d", historial.Id)

	historialEstadoSabatico := models.HistorialEstadoSabatico{
		Id:                historial.Id,
		TerceroId:         historial.TerceroId,
		Justificacion:     historial.Justificacion,
		Activo:            historial.Activo,
		FechaCreacion:     historial.FechaCreacion,
		FechaModificacion: historial.FechaModificacion,
		EstadoSabaticoId:  historial.EstadoSabaticoId,
		SabaticoId:        historial.SabaticoId,
	}

	if err := request.SendJson(
		url,
		"PUT",
		&response,
		historialEstadoSabatico,
	); err != nil {
		return fmt.Errorf(
			"error desactivando historial_estado_sabatico: %v",
			err,
		)
	}

	if err := helpers.ValidateServiceResponse(response); err != nil {
		return fmt.Errorf(
			"sabaticosService historial_estado_sabatico returned error: %w",
			err,
		)
	}

	return nil
}

func ActualizarSoporteSabatico(soporteSabatico models.SoporteSabatico) (*models.SoporteSabatico, error) {
	var response interface{}
	var soporteSabaticoRes models.SoporteSabatico

	url := strings.TrimRight(
		beego.AppConfig.String("sabaticosService"),
		"/",
	) + fmt.Sprintf("/soporte_sabatico/%d", soporteSabatico.Id)

	if err := request.SendJson(
		url,
		"PUT",
		&response,
		soporteSabatico,
	); err != nil {
		return nil, fmt.Errorf(
			"error actualizando soporte_sabatico: %v",
			err,
		)
	}
	if err := helpers.ExtractDataApi(
		response,
		&soporteSabaticoRes,
	); err != nil {
		return nil, fmt.Errorf(
			"error extrayendo soporte_sabatico: %v",
			err,
		)
	}

	return &soporteSabaticoRes, nil
}
