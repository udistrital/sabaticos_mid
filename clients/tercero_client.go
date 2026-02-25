package clients

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func ValidarTercero(terceroId int) error {
	var tercero interface{}
	if err := request.GetJson(beego.AppConfig.String("terceroService")+"tercero/"+fmt.Sprintf("%d", terceroId), &tercero); err != nil {
		beego.Error("Error GET tercero:", err)
		return err
	}

	return nil
}
