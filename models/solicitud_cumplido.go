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
	// _, err = o.Insert(CumplidoProveedor{
	// 	NumeroContrato:   SolicitudCumplido.NumeroContrato,
	// 	VigenciaContrato: SolicitudCumplido.VigenciaContrato,
	// 	Activo:           true,
	// })
	if res, err := GetAllEstadoCumplido(map[string]string{"codigo_abreviacion": "CD"}, []string{}, []string{}, []string{}, 0, 1); err == nil {
		fmt.Println(res)
		//o.Insert(CambioEstadoCumplido{
		//	EstadoCumplidoId: res[0]["Id"].(int),
		//})
	}

	err = o.Commit()

	return

}
