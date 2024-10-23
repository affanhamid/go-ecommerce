package routes

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/controllers"
	"github.com/affanhamid/go-ecommerce/middleware"
)

func LoadRoutes() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/database/create", middleware.CheckMethod("GET", controllers.CreateTablesHandler))
	http.HandleFunc("/database/dropAll", middleware.CheckMethod("GET", controllers.DropTablesHandler))
}
