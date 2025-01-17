package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/udistrital/revision_cumplidos_proveedores_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// TrCrearSolicitudCumplidoController operations for Tr_crear_solicitud_cumplido
type TrCrearSolicitudCumplidoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrCrearSolicitudCumplidoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create SolicitudCumplido
// @Param	body		body 	models.SolicitudCumplido		"body for SolicitudCumplido"
// @Success 201 {object} models.SolicitudCumplido
// @Failure 403 body is empty
// @router / [post]
func (c *TrCrearSolicitudCumplidoController) Post() {

	var v models.SolicitudCumplido

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if err = models.CrearSolicitudCumplido(&v); err == nil {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful modification", "Data": v}

		} else {
			logs.Error(err)
			c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
			c.Abort("400")
		}
	} else {
		fmt.Println(err)
		logs.Error(err)
		c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}

// // GetOne ...
// // @Title GetOne
// // @Description get Tr_crear_solicitud_cumplido by id
// // @Param	id		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.Tr_crear_solicitud_cumplido
// // @Failure 403 :id is empty
// // @router /:id [get]
// func (c *TrCrearSolicitudCumplidoController) GetOne() {

// }

// // GetAll ...
// // @Title GetAll
// // @Description get Tr_crear_solicitud_cumplido
// // @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// // @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// // @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// // @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// // @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// // @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// // @Success 200 {object} models.Tr_crear_solicitud_cumplido
// // @Failure 403
// // @router / [get]
// func (c *TrCrearSolicitudCumplidoController) GetAll() {

// }

// // Put ...
// // @Title Put
// // @Description update the Tr_crear_solicitud_cumplido
// // @Param	id		path 	string	true		"The id you want to update"
// // @Param	body		body 	models.Tr_crear_solicitud_cumplido	true		"body for Tr_crear_solicitud_cumplido content"
// // @Success 200 {object} models.Tr_crear_solicitud_cumplido
// // @Failure 403 :id is not int
// // @router /:id [put]
// func (c *TrCrearSolicitudCumplidoController) Put() {

// }

// // Delete ...
// // @Title Delete
// // @Description delete the Tr_crear_solicitud_cumplido
// // @Param	id		path 	string	true		"The id you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 id is empty
// // @router /:id [delete]
// func (c *TrCrearSolicitudCumplidoController) Delete() {

// }
