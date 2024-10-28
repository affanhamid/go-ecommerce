package routes

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/app"
	"github.com/affanhamid/go-ecommerce/middleware"
)

func LoadUserRoutes(app *app.App) {
	// user
	http.HandleFunc("/api/add-user", middleware.CheckMethod("POST", app.Controller.CreateUserHandler))
	http.HandleFunc("/api/user", middleware.CheckMethod("GET", app.Controller.GetUserHandler))
	http.HandleFunc("/api/delete-user", middleware.CheckMethod("POST", app.Controller.DeleteUserHandler))
	http.HandleFunc("/api/delete-user-permanent", middleware.CheckMethod("POST", app.Controller.DeleteUserPermaHandler))

	// address
	http.HandleFunc("/api/add-address", middleware.CheckMethod("POST", app.Controller.AddAddressHandler))
	http.HandleFunc("/api/delete-address", middleware.CheckMethod("POST", app.Controller.DeleteAddressHandler))
}
