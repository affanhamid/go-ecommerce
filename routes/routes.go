package routes

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/controllers"
	"github.com/affanhamid/go-ecommerce/database"
	"github.com/affanhamid/go-ecommerce/middleware"
)

func LoadRoutes() {
	db := database.Connect()

	app := controllers.App{DB: db}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./static/signup.html") })
	http.HandleFunc("/delete-account", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./static/delete.html") })

	http.HandleFunc("/api/admin/create", middleware.CheckMethod("POST", app.CreateTablesHandler))
	http.HandleFunc("/api/admin/drop-all", middleware.CheckMethod("POST", app.DropTablesHandler))
	http.HandleFunc("/api/add-user", middleware.CheckMethod("POST", app.CreateUserHandler))
	http.HandleFunc("/api/delete-user", middleware.CheckMethod("POST", app.DeleteUserHandler))
}
