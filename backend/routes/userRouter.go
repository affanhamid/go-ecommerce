package routes

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/app"
	"github.com/affanhamid/go-ecommerce/middleware"
)

func LoadUserRoutes(app *app.App) {
	http.HandleFunc("/api/add-user", middleware.CheckMethod("POST", app.Controller.CreateUserHandler))
	http.HandleFunc("/api/delete-user", middleware.CheckMethod("POST", app.Controller.DeleteUserHandler))
	http.HandleFunc("/api/delete-user-permanent", middleware.CheckMethod("POST", app.Controller.DeleteUserPermaHandler))
}
