package controllers

import (
	"net/http"

	"github.com/udistrital/sabaticos_mid/helpers"
	"github.com/udistrital/sabaticos_mid/models"
	"github.com/udistrital/sabaticos_mid/service"

	"github.com/astaxie/beego"
)

// SoporteSolicitudController operations for SoporteSolicitud
type SoporteSolicitudController struct {
	beego.Controller
}

// URLMapping ...
func (c *SoporteSolicitudController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create SoporteSolicitud con múltiples documentos
// @router / [post]
func (c *SoporteSolicitudController) Post() {
	var soporteSolicitudRequest models.SoporteSolicitudRequest

	// Obtener datos de form-data
	soporteSolicitudRequest.TerceroId, _ = c.GetInt("tercero_id")
	soporteSolicitudRequest.SolicitudId, _ = c.GetInt("solicitud_id")
	soporteSolicitudRequest.RolUsuario = c.GetString("rol_usuario")
	soporteSolicitudRequest.EstadoSoporteSolicitud = c.GetString("estado_soporte_solicitud")
	soporteSolicitudRequest.TipoDocumentoId, _ = c.GetInt("tipo_documento_id")

	// Validar campos requeridos
	if soporteSolicitudRequest.TerceroId == 0 || soporteSolicitudRequest.SolicitudId == 0 ||
		soporteSolicitudRequest.RolUsuario == "" || soporteSolicitudRequest.EstadoSoporteSolicitud == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "fields terceroId, solicitudId, estadoSoporteSolicitud and rolUsuario are required")
		return
	}

	// Obtener archivos
	files, _ := c.GetFiles("documentos")

	if len(files) == 0 {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "at least one file must be provided")
		return
	}

	respuesta, err := service.CrearSoporteSolicitud(soporteSolicitudRequest, files[0])

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, err.Error())
		return
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, respuesta, "")
}
