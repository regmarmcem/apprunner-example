package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/regmarmcem/apprunner-example/controllers"
	"github.com/regmarmcem/apprunner-example/services"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	ser := services.NewTaskAppService(db)
	tCon := controllers.NewTaskController(ser)

	r.MethodFunc(http.MethodGet, "/task", tCon.GetTaskHandler)

	return r
}
