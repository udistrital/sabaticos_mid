package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/requestresponse"
)

func ExtractData(response interface{}) interface{} {

	// Validar si es map dinámico
	if respMap, ok := response.(map[string]interface{}); ok {
		// Soporta distintos envelopes: Data/data/res/result/Result
		for _, key := range []string{"Data", "data", "res", "result", "Result"} {
			if data, exists := respMap[key]; exists {
				if data == nil {
					if errVal, hasErr := respMap["Error"]; hasErr {
						return errVal
					}
				}
				return data
			}
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

// ValidateServiceResponse detecta respuestas de error comunes de servicios OAS
func ValidateServiceResponse(response interface{}) error {
	respMap, ok := response.(map[string]interface{})
	if !ok {
		return nil
	}

	if success, exists := respMap["Success"]; exists {
		if okSuccess, ok := success.(bool); ok && !okSuccess {
			msg, _ := respMap["Message"].(string)
			if msg == "" {
				msg, _ = respMap["Status"].(string)
			}
			if msg == "" {
				msg = "service response with Success=false"
			}
			return errors.New(msg)
		}
	}

	code := firstStatusCode(respMap)
	if code >= 400 {
		status, _ := respMap["Status"].(string)
		message, _ := respMap["Message"].(string)
		errorText, _ := respMap["Error"].(string)

		detail := message
		if detail == "" {
			detail = errorText
		}
		if detail == "" {
			detail = status
		}
		if detail == "" {
			detail = "unspecified error"
		}

		return fmt.Errorf("code %d: %s", code, detail)
	}

	return nil
}

func firstStatusCode(respMap map[string]interface{}) int {
	keys := []string{"Code", "code", "StatusCode", "statusCode", "Status"}
	for _, key := range keys {
		if raw, ok := respMap[key]; ok {
			if code := toInt(raw); code > 0 {
				return code
			}
		}
	}
	return 0
}

func toInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		n, err := strconv.Atoi(v)
		if err == nil {
			return n
		}
		// Supports formats like "500 Internal Server Error"
		for i := 0; i < len(v); i++ {
			if v[i] < '0' || v[i] > '9' {
				if i > 0 {
					n, err = strconv.Atoi(v[:i])
					if err == nil {
						return n
					}
				}
				break
			}
		}
	}
	return 0
}

// JSONResponse envía una respuesta estándar en formato APIResponseDTO
func JSONResponse(
	c *beego.Controller,
	success bool,
	status int,
	data interface{},
	message string,
) {
	data = ExtractData(data) // Cleans OAS-type data if necessary

	// Set standard JSON response
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = requestresponse.APIResponseDTO(
		success,
		status,
		data,
		message,
	)
	c.ServeJSON()
}

// JoinStrings joins a slice of strings with a separator
func JoinStrings(strings []string, sep string) string {
	result := ""
	for i, s := range strings {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
