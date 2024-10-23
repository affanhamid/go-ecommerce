package routes

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/controllers"
	"github.com/affanhamid/go-ecommerce/middleware"
)

func LoadRoutes() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/admin/create", middleware.CheckMethod("POST", controllers.CreateTablesHandler))
	http.HandleFunc("/api/admin/drop-all", middleware.CheckMethod("POST", controllers.DropTablesHandler))
	http.HandleFunc("/api/add-user", middleware.CheckMethod("POST", controllers.AddUser))
}
