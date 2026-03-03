package controllers

import (
	"encoding/json"
	"net/http"

	"api_mid_sabaticos/clients"
	"api_mid_sabaticos/helpers"
	"api_mid_sabaticos/models"
	"api_mid_sabaticos/service"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestmanager"
)

// SolicitudController operations for Solicitud
// @Tag Solicitud
type SolicitudController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Solicitud
// @Param	body		body 	models.Solicitud	true		"body for Solicitud content"
// @Success 201 {object} models.Solicitud
// @Failure 403 body is empty
// @router / [post]
func (c *SolicitudController) Post() {
	defer errorhandler.HandlePanic(&c.Controller)

	var solicitudRequest models.SolicitudRequest

	requestmanager.FillRequestWithPanic(&c.Controller, &solicitudRequest)

	if solicitudRequest.TerceroId <= 0 || solicitudRequest.TipoSolicitudId == nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "Los campos terceroId y tipoSolicitudId son requeridos")
	}

	solicitud, historico, formulario, err := service.CrearSolicitud(solicitudRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadGateway, nil, "Error al registrar solicitud")
	}

	respuesta := models.SolicitudResponse{
		Solicitud:  solicitud,
		Historial:  historico,
		Formulario: formulario,
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusCreated, respuesta, "Solicitud creada exitosamente")
}

// GetOne ...
// @Title GetOne
// @Description get Solicitud by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Solicitud
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SolicitudController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Solicitud
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Solicitud
// @Failure 403
// @router / [get]
func (c *SolicitudController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Solicitud
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Solicitud	true		"body for Solicitud content"
// @Success 200 {object} models.Solicitud
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SolicitudController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Solicitud
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SolicitudController) Delete() {

}

func (c *SolicitudController) Validar() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.Ctx.Input.CopyBody(1 << 20) // Beego v1 [web:30][web:64]

	var req models.ValidarSolicitudRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "JSON inválido: "+err.Error())
		return
	}

	if req.TerceroId <= 0 || req.SolicitudId <= 0 {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "Los campos tercero_id y solicitud_id son requeridos")
		return
	}

	// Armar objetos que tu client ya espera (sin romper compatibilidad)
	solicitudCreada := &models.Solicitud{Id: req.SolicitudId}
	solicitudReq := models.SolicitudRequest{TerceroId: req.TerceroId}

	// Overrides opcionales: si vienen en el JSON, tu client los toma; si no, usa defaults
	historialOverride := models.HistorialSolicitud{
		EstadoSolicitudId: req.EstadoSolicitudId,
		Justificacion:     req.Justificacion,
	}

	historialCreado, err := clients.RegistrarHistorialSolicitud(historialOverride, solicitudCreada, solicitudReq)
	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadGateway, nil, "Error al registrar historial: "+err.Error())
		return
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusCreated, historialCreado, "Historial registrado exitosamente")
}
