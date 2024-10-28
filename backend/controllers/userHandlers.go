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

type Controller struct {
	DB *gorm.DB
}

func (c *Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var user models.User
	if !utils.ValidateAndDecode(w, r, &user) {
		return
	}

	if err := database.CreateUser(&user, c.DB); err != nil {

		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			switch pgError.Code {
			case "23505": // Unique Violation
				utils.SendJSONMessage(w, http.StatusConflict, "User already exists")
			default:
				utils.SendJSONMessage(w, http.StatusInternalServerError, "Database Error"+pgError.Message)
			}
		} else {
			// Non-Postgres error
			utils.SendJSONMessage(w, http.StatusInternalServerError, "Failed to add user: "+err.Error())
		}
		return
	}
	utils.SendJSONMessage(w, http.StatusOK, "User data added successfully!")
}

func (c *Controller) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var userEmail models.UserEmail
	if !utils.ValidateAndDecode(w, r, &userEmail) {
		return
	}

	user, err := database.GetUser(userEmail.Email, c.DB)
	if err != nil {
		utils.SendJSONMessage(w, http.StatusInternalServerError, "Failed to fetch user: "+err.Error())
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error conveerting user to JSON object: "+err.Error(), http.StatusInternalServerError)
	}

	utils.SendJSON(w, http.StatusOK, userJSON)
}

func (c *Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var userEmail models.UserEmail
	if !utils.ValidateAndDecode(w, r, &userEmail) {
		return
	}

	if err := database.DeleteUser(&userEmail, c.DB); err != nil {
		utils.SendJSONMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.SendJSONMessage(w, http.StatusOK, "User deleted succesfully!")
}

func (c *Controller) DeleteUserPermaHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var userEmail models.UserEmail
	if !utils.ValidateAndDecode(w, r, &userEmail) {
		return
	}

	if err := database.DeleteUserPerma(&userEmail, c.DB); err != nil {
		utils.SendJSONMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.SendJSONMessage(w, http.StatusOK, "User deleted permanently succesfully!")
}
