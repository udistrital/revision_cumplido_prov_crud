package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCambioEstadoCumplido(t *testing.T) {
	convey.Convey("Test: /v1/cambio_estado_cumplido/ - GetAll cambio_estado_cumplido", t, func() {
		r, err := http.NewRequest("GET", "/v1/cambio_estado_cumplido/", nil)
		if err != nil {
			t.Fatal("Error al crear la solicitud:", err)
		}
		w := httptest.NewRecorder()

		beego.BeeApp.Handlers.ServeHTTP(w, r)

		convey.Convey("Status Code Should Be 200", func() {
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
		})

		convey.Convey("The Result Should Not Be Empty", func() {
			convey.So(w.Body.Len(), convey.ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetOneCambioEstadoCumplidoId(t *testing.T) {
	convey.Convey("Test: /v1/cambio_estado_cumplido/ - GetOne cambio_estado_cumplido", t, func() {
		r, err := http.NewRequest("GET", "/v1/cambio_estado_cumplido/5", nil)
		if err != nil {
			t.Fatal("Error al crear la solicitud:", err)
		}
		w := httptest.NewRecorder()

		beego.BeeApp.Handlers.ServeHTTP(w, r)

		convey.Convey("Status Code Should Be 200", func() {
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
		})

		convey.Convey("The Result Should Not Be Empty", func() {
			convey.So(w.Body.Len(), convey.ShouldBeGreaterThan, 0)
		})
	})
}
