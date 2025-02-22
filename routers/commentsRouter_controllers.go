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

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:InformacionPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TrCrearSolicitudCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/revision_cumplidos_proveedores_crud/controllers:TrCrearSolicitudCumplidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
