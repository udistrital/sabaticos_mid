package controllers

import (
	"net/http"

	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"

	"github.com/astaxie/beego"
)

// SoporteSolicitudController operations for SoporteSolicitud
type SoporteSabaticoController struct {
	beego.Controller
}

// URLMapping ...
func (c *SoporteSabaticoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create SoporteSabatico con múltiples documentos
// @router / [post]
func (c *SoporteSabaticoController) Post() {
	var soporteSolicitudRequest models.SoporteSabatcioRequest

	// Obtener datos de form-data
	soporteSolicitudRequest.SabaticoId, _ = c.GetInt("SabaticoId")
	soporteSolicitudRequest.RolUsuario = c.GetString("rol_usuario")
	soporteSolicitudRequest.EstadoSoporteSabatico = c.GetString("estado_soporte_sabatico")

	// Validar campos requeridos
	if soporteSolicitudRequest.SabaticoId == 0 ||
		soporteSolicitudRequest.RolUsuario == "" || soporteSolicitudRequest.EstadoSoporteSabatico == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "fields terceroId, SabaticoId, estadoSoporteSolicitud and rolUsuario are required")
		return
	}

	// Obtener archivos
	files, _ := c.GetFiles("documentos")

	if len(files) == 0 {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "at least one file must be provided")
		return
	}

	respuesta, err := service.CrearSoporteSabatico(soporteSolicitudRequest, files[0])

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, err.Error())
		return
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, respuesta, "")
}
