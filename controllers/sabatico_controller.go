package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestmanager"
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
// @router / [post]
func (c *SabaticoController) PostCrearSabatico() {
	defer errorhandler.HandlePanic(&c.Controller)
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

	result, err := service.CrearSabatico(
		req.SolicitudId,
		req.TerceroId,
		req.Observaciones,
		req.FechaInicio,
		req.FechaFin,
	)

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

// PostGuardarPlanTrabajoSabatico ...
// @Title PostGuardarPlanTrabajoSabatico
// @Description Actualiza el plan de trabajo sabático desactivando el historial actual y creando un nuevo registro con la información actualizada
// @Param body body models.PlanTrabajoSabaticoRequest true "Body para actualizar el plan de trabajo sabático (descripcion e historial_estado_sabatico_id)"
// @Success 201 {object} models.HistorialEstadoSabatico
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @router /plan_trabajo [post]
func (c *SabaticoController) PostGuardarPlanTrabajoSabatico() {
	defer errorhandler.HandlePanic(&c.Controller)
	var planTrabajoSabaticoRequest models.PlanTrabajoSabaticoRequest

	requestmanager.FillRequestWithPanic(&c.Controller, &planTrabajoSabaticoRequest)

	if planTrabajoSabaticoRequest.Justificacion == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "The camp justificacion is required")
	}

	result, err := service.GuardarPlanTrabajoSabatico(planTrabajoSabaticoRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "error filing request: "+err.Error())
		return
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, result, "request filed successfully")

}

// CambiarEstadoPlanTrabajoSabatico ...
// @Title CambiarEstadoPlanTrabajoSabatico
// @Description Cambia el estado del plan de trabajo sabático y actualiza el estado del soporte asociado
// @Param body body models.AprobarRechazarPlanTRabajoSabaticoequest true "Body para cambiar el estado del plan de trabajo sabático"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @router /plan_trabajo/estado [post]
func (c *SabaticoController) CambiarEstadoPlanTrabajoSabatico() {
	defer errorhandler.HandlePanic(&c.Controller)

	var cambiarEstadoPlanTrabajoRequest models.AprobarRechazarPlanTRabajoSabaticoequest

	requestmanager.FillRequestWithPanic(&c.Controller, &cambiarEstadoPlanTrabajoRequest)

	if cambiarEstadoPlanTrabajoRequest.SabaticoId <= 0 {
		helpers.JSONResponse(
			&c.Controller,
			false,
			http.StatusBadRequest,
			nil,
			"The Sabbatical field is necessary",
		)
		return
	}

	if cambiarEstadoPlanTrabajoRequest.EstadoSabatico == "" {
		helpers.JSONResponse(
			&c.Controller,
			false,
			http.StatusBadRequest,
			nil,
			"The field EstadoSabatico is required",
		)
		return
	}

	if cambiarEstadoPlanTrabajoRequest.EstadoSoporteSabatico == "" {
		helpers.JSONResponse(
			&c.Controller,
			false,
			http.StatusBadRequest,
			nil,
			"The SabbaticalSupportState field is required",
		)
		return
	}

	result, err := service.CambiarEstadoPlanTrabajoSabatico(cambiarEstadoPlanTrabajoRequest)
	if err != nil {
		helpers.JSONResponse(
			&c.Controller,
			false,
			http.StatusInternalServerError,
			nil,
			"Error changing request status: "+err.Error(),
		)
		return
	}

	helpers.JSONResponse(
		&c.Controller,
		true,
		http.StatusOK,
		result,
		"Application processed successfully",
	)
}
