package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sabaticos_mid/clients"
	"github.com/udistrital/sabaticos_mid/models"
)

func CrearSabatico(terceroId int, observaciones string, fechaInicio string, fechaFin string, estado string) (*models.CrearSabaticoResult, error) {
	fmt.Println("--------------------- Entra a Registrar Sabatico ---------------------")

	estadoSabaticoId, err := clients.ConsultarIdEstadoSabatico(estado)
	if err != nil {
		return nil, fmt.Errorf("error consultando id de estado sabático: %v", err)
	}

	crudURL := strings.TrimRight(beego.AppConfig.String("sabaticosService"), "/") + "/sabatico"

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
