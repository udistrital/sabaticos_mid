package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"
)

type SabaticoController struct {
	beego.Controller
}

// PostCrearSabatico ...
// @Title PostCrearSabatico
// @Description Crea un sabático consumiendo el CRUD
// @Param	body	body  models.CrearSabaticoRequest	true	"Body para crear sabático"
// @Success 201 {object} models.CrearSabaticoResponse
// @Failure 400 {object} models.CrearSabaticoResponse
// @Failure 500 {object} models.CrearSabaticoResponse
// @router /crear [post]
func (c *SabaticoController) PostCrearSabatico() {
	fmt.Println("-------------- Entra al Controller ---------------")
	var req models.CrearSabaticoRequest

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = models.CrearSabaticoResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "body inválido",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	result, err := service.CrearSabatico(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = models.CrearSabaticoResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = models.CrearSabaticoResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "sabático creado correctamente",
		Data:    result,
	}
	c.ServeJSON()
}
