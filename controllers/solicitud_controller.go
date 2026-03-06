package controllers

import (
	"net/http"

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
	c.Mapping("Radicar", c.Radicar)
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

	if solicitudRequest.TerceroId <= 0 || solicitudRequest.TipoSolicitudId == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "Los campos terceroId y tipoSolicitudId son requeridos")
	}

	solicitud, err := service.CrearSolicitud(solicitudRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "Recurso no encontrado: "+err.Error())
		return
	}

	respuesta := models.SolicitudResponse{
		Solicitud: solicitud,
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

// Radicar ...
// @Title Radicar
// @Description Radicar una solicitud (cambiar estado y crear registros)
// @Param	id		path 	int	true		"The id of the Solicitud to radicar"
// @Param	body		body 	models.SolicitudRequest	false	"body for additional data if needed"
// @Success 200 {object} models.SolicitudResponse
// @Failure 400 Bad request
// @Failure 404 Solicitud not found
// @router /radicar/:id [post]
func (c *SolicitudController) Radicar() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.GetString(":id")
	var RadicarSolicitudRequest models.RadicarSolicitudRequest

	requestmanager.FillRequestWithPanic(&c.Controller, &RadicarSolicitudRequest)

	// Validar que el ID de la solicitud esté presente
	if id == "" {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "El id de la solicitud es requerido")
		return
	}

	//Validar RadicarSolicitudRequest si es necesario, por ejemplo:
	if RadicarSolicitudRequest.SolicitudId == 0 || RadicarSolicitudRequest.FormularioId == 0 || RadicarSolicitudRequest.Formulario == nil || RadicarSolicitudRequest.DocumentosId == nil || len(RadicarSolicitudRequest.DocumentosId) == 0 {
		helpers.JSONResponse(&c.Controller, false, http.StatusBadRequest, nil, "Los campos SolicitudId, FormularioId, Formulario y DocumentosId son requeridos en el cuerpo de la solicitud")
		return
	}

	// Llamar al servicio para radicar
	solicitud, err := service.RadicarSolicitud(RadicarSolicitudRequest)

	if err != nil {
		helpers.JSONResponse(&c.Controller, false, http.StatusNotFound, nil, "Error al radicar solicitud: "+err.Error())
		return
	}

	respuesta := models.SolicitudResponse{
		Solicitud: solicitud,
	}

	helpers.JSONResponse(&c.Controller, true, http.StatusOK, respuesta, "Solicitud radicada exitosamente")
}
