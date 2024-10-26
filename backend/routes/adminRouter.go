package routes

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/app"
	"github.com/affanhamid/go-ecommerce/middleware"
)

func LoadAdminRoutes(app *app.App) {
	http.HandleFunc("/api/admin/create", middleware.CheckMethod("POST", app.Controller.CreateTablesHandler))
	http.HandleFunc("/api/admin/drop-all", middleware.CheckMethod("POST", app.Controller.DropTablesHandler))
}
