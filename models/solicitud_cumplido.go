package models

import (
	"encoding/json"
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

	fmt.Println(cumplido_proveedor)
	id_cumplido_proveedor, err := o.Insert(&cumplido_proveedor)

	if err != nil {
		o.Rollback()
		return
	}

	if res, err := GetAllEstadoCumplido(map[string]string{"codigo_abreviacion": "CD"}, []string{}, []string{}, []string{}, 0, 1); err == nil {
		//fmt.Println(res)
		fmt.Println("interface", res[0])
		b := []byte(`{"key":"value"}`)
		json.Unmarshal(b, &res[0])
		ec := res[0].(map[string]interface{})
		fmt.Println(ec)
		_, err = o.Insert(
			CambioEstadoCumplido{
				EstadoCumplidoId: &EstadoCumplido{
					Id: ec["Id"].(int),
				},
				CumplidoProveedorId: &CumplidoProveedor{
					Id: int(id_cumplido_proveedor),
				},
				DocumentoResponsable: m.DocumentoResponsable,
				CargoResponsable:     m.CargoResponsable,
			})
		if err != nil {
			o.Rollback()
		}
	} else {
		o.Rollback()

	}

	err = o.Commit()

	return

}
