package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/smartystreets/goconvey/convey"
	"github.com/udistrital/revision_cumplidos_proveedores_crud/controllers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	pgUser := os.Getenv("REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGUSER")
	pgPass := os.Getenv("REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGPASS")
	pgHost := os.Getenv("REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGHOST")
	pgPort := os.Getenv("REVISION_CUMPLIDOS_PROOVEDORES_CRUD_GPORT")
	pgDb := os.Getenv("REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGDB")
	pgSchema := os.Getenv("REVISION_CUMPLIDOS_PROOVEDORES_CRUD_PGSCHEMA")

	connStr := "postgres://" + pgUser + ":" + pgPass + "@" + pgHost + ":" + pgPort + "/" + pgDb + "?sslmode=disable&search_path=" + pgSchema

	orm.RegisterDataBase("default", "postgres", connStr)
	beego.BeeApp = beego.NewApp()
	beego.Router("/v1/estado_cumplido/", &controllers.EstadoCumplidoController{}, "get:GetAll")
	beego.Router("/v1/estado_cumplido/:id", &controllers.EstadoCumplidoController{}, "get:GetOne")
	beego.Router("/v1/cambio_estado_cumplido/", &controllers.CambioEstadoCumplidoController{}, "get:GetAll")
	beego.Router("/v1/cambio_estado_cumplido/:id", &controllers.CambioEstadoCumplidoController{}, "get:GetOne")

}

func TestGetAllEstadoCumplido(t *testing.T) {
	convey.Convey("Test: /v1/estado_cumplido/ - GetAll estado_cumplido", t, func() {
		r, err := http.NewRequest("GET", "/v1/estado_cumplido/", nil)
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

func TestGetOneEstadoCumplidoId(t *testing.T) {
	convey.Convey("Test: /v1/estado_cumplido/ - GetOne estado_cumplido", t, func() {
		r, err := http.NewRequest("GET", "/v1/estado_cumplido/5", nil)
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

func TestGetAllEstadoCumplido_Fail(t *testing.T) {
	convey.Convey("Test: /v1/estado_cumplido/ - GetAll estado_cumplido - Failing Case", t, func() {
		r, err := http.NewRequest("GET", "/v1/estado_cumplido/invalid", nil) // Ruta incorrecta
		if err != nil {
			t.Fatal("Error al crear la solicitud:", err)
		}
		w := httptest.NewRecorder()

		beego.BeeApp.Handlers.ServeHTTP(w, r)

		convey.Convey("Status Code Should Be 400", func() {
			convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest) // Este código debería fallar
		})

		convey.Convey("The Result Should Not Be Empty", func() {
			convey.So(w.Body.Len(), convey.ShouldBeGreaterThan, 0) // Este código también debería fallar
		})
	})
}
