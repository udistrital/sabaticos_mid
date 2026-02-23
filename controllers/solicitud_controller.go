package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api_mid_sabaticos/models"
	"api_mid_sabaticos/service"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/requestresponse"
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
	var solicitud models.SolicitudRequest

	fmt.Println(c.Ctx.Input.RequestBody)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &solicitud); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = requestresponse.APIResponseDTO(false, http.StatusBadRequest, nil, "Formato de solicitud inválido")
		c.ServeJSON()
		return
	}

	if solicitud.TerceroId <= 0 && solicitud.TipoSolicitudId == nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = requestresponse.APIResponseDTO(false, http.StatusBadRequest, nil, "El campo terceroId es requerido")
		c.ServeJSON()
		return
	}

	respuesta, err := service.StoreSolicitud(solicitud)

	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadGateway)
		c.Data["json"] = requestresponse.APIResponseDTO(false, http.StatusBadGateway, nil, "Error al registrar solicitud en el CRUD")
		c.ServeJSON()
		return
	}

	c.Data["json"] = respuesta
	c.ServeJSON()
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
