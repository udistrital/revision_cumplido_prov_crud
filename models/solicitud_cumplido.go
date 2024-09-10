package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type SolicitudCumplido struct {
	NumeroContrato       string
	VigenciaContrato     int
	DocumentoResponsable int
	CargoResponsable     string
}

// transaction create a cumplido_provedor con cambio_estado_inicial de cargando_documentos
func CrearSolicitudCumplido(m *SolicitudCumplido) (err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		fmt.Println(err)
	}
	var cumplido_proveedor CumplidoProveedor

	cumplido_proveedor.NumeroContrato = m.NumeroContrato
	cumplido_proveedor.VigenciaContrato = m.VigenciaContrato
	cumplido_proveedor.Activo = true

	id_cumplido_proveedor, err := o.Insert(&cumplido_proveedor)

	if err != nil {
		o.Rollback()
		return
	}

	if res, err := GetAllEstadoCumplido(map[string]string{"codigo_abreviacion": "CD"}, []string{}, []string{}, []string{}, 0, 1); err == nil {

		var cambio_estado_cumplido CambioEstadoCumplido
		var estado_cumplido EstadoCumplido
		estado_cumplido.Id = res[0].(EstadoCumplido).Id
		cumplido_proveedor.Id = int(id_cumplido_proveedor)

		cambio_estado_cumplido.EstadoCumplidoId = &estado_cumplido
		cambio_estado_cumplido.CumplidoProveedorId = &cumplido_proveedor
		cambio_estado_cumplido.DocumentoResponsable = m.DocumentoResponsable
		cambio_estado_cumplido.CargoResponsable = m.CargoResponsable
		cambio_estado_cumplido.Activo = true

		if _, err = o.Insert(&cambio_estado_cumplido); err != nil {
			o.Rollback()
		}
	} else {
		o.Rollback()
	}

	err = o.Commit()

	return

}
