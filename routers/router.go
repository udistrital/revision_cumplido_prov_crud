// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/revision_cumplidos_proveedores_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/cumplido_proveedor",
			beego.NSInclude(
				&controllers.CumplidoProveedorController{},
			),
		),

		beego.NSNamespace("/soporte_cumplido",
			beego.NSInclude(
				&controllers.SoporteCumplidoController{},
			),
		),

		beego.NSNamespace("/cambio_estado_cumplido",
			beego.NSInclude(
				&controllers.CambioEstadoCumplidoController{},
			),
		),

		beego.NSNamespace("/estado_cumplido",
			beego.NSInclude(
				&controllers.EstadoCumplidoController{},
			),
		),

		beego.NSNamespace("/comentario_soporte",
			beego.NSInclude(
				&controllers.ComentarioSoporteController{},
			),
		),

		beego.NSNamespace("/informacion_pago",
			beego.NSInclude(
				&controllers.InformacionPagoController{},
			),
		),
		beego.NSNamespace("/tipo_pago",
			beego.NSInclude(
				&controllers.TipoPagoController{},
			),
		),
		beego.NSNamespace("/crear_solicitud_cumplido",
			beego.NSInclude(
				&controllers.TrCrearSolicitudCumplidoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
