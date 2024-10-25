package routes

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/controllers"
	"github.com/affanhamid/go-ecommerce/database"
	"github.com/affanhamid/go-ecommerce/middleware"
)

func LoadUserRoutes() {
	db := database.Connect()
	app := controllers.App{DB: db}

	http.HandleFunc("/api/add-user", middleware.CheckMethod("POST", app.CreateUserHandler))
	http.HandleFunc("/api/delete-user", middleware.CheckMethod("POST", app.DeleteUserHandler))
	http.HandleFunc("/api/delete-user-permanent", middleware.CheckMethod("POST", app.DeleteUserPermaHandler))
}
