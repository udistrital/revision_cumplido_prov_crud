package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoSoporteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoportePagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
