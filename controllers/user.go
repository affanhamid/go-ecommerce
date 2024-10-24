package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/affanhamid/go-ecommerce/database"
	"github.com/affanhamid/go-ecommerce/models"
	"github.com/affanhamid/go-ecommerce/utils"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (app *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var user models.User
	if !utils.ValidateAndDecode(w, r, &user) {
		return
	}

	if err := database.CreateUser(&user, app.DB); err != nil {

		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			switch pgError.Code {
			case "23505": // Unique Violation
				utils.SendJSONError(w, http.StatusConflict, "User already exists with that email")
			default:
				utils.SendJSONError(w, http.StatusInternalServerError, "Database Error"+pgError.Message)
			}
		} else {
			// Non-Postgres error
			utils.SendJSONError(w, http.StatusInternalServerError, "Failed to add user: "+err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User data added successfully!"})
}

func (app *App) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var user models.DeleteUser
	if !utils.ValidateAndDecode(w, r, &user) {
		return
	}

	if err := database.DeleteUser(&user, app.DB); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully!"})
}
