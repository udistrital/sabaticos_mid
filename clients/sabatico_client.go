package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

func ConsultarSabatico(sabaticoId int) (*models.Sabatico, error) {
	fmt.Println("--------------------- Entra a Consultar Sabatico ---------------------")
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

func RegistrarSabatico(terceroId int, observaciones string, fechaInicio string, fechaFin string, estadoSabaticoId int) (*models.CrearSabaticoResult, error) {
	fmt.Println("--------------------- Entra a Registrar Sabatico ---------------------")
	crudURL := beego.AppConfig.String("sabaticosService") + "/sabatico"

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

	logs.Info("payload registrar sabatico: %s", string(body))
	logs.Info("url crud registrar sabatico: %s", crudURL)

	req, err := http.NewRequest("POST", crudURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creando request al CRUD: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("error consumiendo sabaticos_crud:", err)
		return nil, fmt.Errorf("error consumiendo sabaticos_crud: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta del CRUD: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("sabaticos_crud respondió con estado %d: %s", resp.StatusCode, string(respBytes))
	}

	var result models.CrearSabaticoResult
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta del CRUD: %v", err)
	}

	return &result, nil
}
