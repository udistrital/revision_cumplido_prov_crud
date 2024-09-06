package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/revision_cumplidos_proveedores_crud/models"
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
	if res, err := models.GetAllEstadoCumplido(map[string]string{"CodigoAbreviacion": "CD"}, []string{}, []string{}, []string{}, 0, 1); err == nil {
		fmt.Println(res)
		// o.Insert(models.CambioEstadoCumplido{
		// 	EstadoCumplidoId: res[0],

		// })
	}

	err = o.Commit()

	return

}
