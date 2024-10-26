package controllers

import (
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

func (c *Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var user models.DeleteUser
	if !utils.ValidateAndDecode(w, r, &user) {
		return
	}

	if err := database.DeleteUser(&user, c.DB); err != nil {
		utils.SendJSONMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.SendJSONMessage(w, http.StatusOK, "User deleted succesfully!")
}

func (c *Controller) DeleteUserPermaHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var user models.DeleteUser
	if !utils.ValidateAndDecode(w, r, &user) {
		return
	}

	if err := database.DeleteUserPerma(&user, c.DB); err != nil {
		utils.SendJSONMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.SendJSONMessage(w, http.StatusOK, "User deleted permanently succesfully!")
}