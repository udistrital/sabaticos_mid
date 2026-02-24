package helpers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/requestresponse"
)

func ExtractData(response interface{}) interface{} {

	// Validar si es map dinámico
	if respMap, ok := response.(map[string]interface{}); ok {

		// Si tiene campo Data, lo extrae; si es nil, usa Error
		if data, exists := respMap["Data"]; exists {
			if data == nil {
				if errVal, hasErr := respMap["Error"]; hasErr {
					return errVal
				}
			}
			return data
		}
	}

	return response
}

// ExtractDataApi extrae el campo Data de una respuesta OAS y lo convierte al tipo destino
func ExtractDataApi(response interface{}, dest interface{}) error {
	data := ExtractData(response)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, dest)
}

// JSONResponse envía una respuesta estándar en formato APIResponseDTO
func JSONResponse(
	c *beego.Controller,
	success bool,
	status int,
	data interface{},
	message string,
) {
	data = ExtractData(data) // Limpia datos tipo OAS si es necesario

	// Configurar respuesta JSON estándar
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = requestresponse.APIResponseDTO(
		success,
		status,
		data,
		message,
	)
	c.ServeJSON()
}
