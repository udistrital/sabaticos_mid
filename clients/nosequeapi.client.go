package clients

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"github.com/udistrital/sabaticos_mid/models"
)

func ConsultarSecretariaAcademicaDocumentoUserId(documentoId string) (*models.Persona, error) {
	// URL del servicio
	url := beego.AppConfig.String("sabaticosService") + "academica_pruebas/secretaria_academica/" + documentoId

	// Crear cliente HTTP con timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Hacer la petición GET
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Validar status code
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code: " + strconv.Itoa(resp.StatusCode))
	}

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parsear XML
	var secretariaResp models.SecretariaAcademicaResponse
	err = xml.Unmarshal(body, &secretariaResp)
	if err != nil {
		return nil, err
	}

	// Validar que se obtuvo la persona
	if secretariaResp.Persona == nil {
		return nil, errors.New("no persona found in response")
	}

	return secretariaResp.Persona, nil
}
