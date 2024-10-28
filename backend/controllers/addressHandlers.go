package controllers

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/database"
	"github.com/affanhamid/go-ecommerce/models"
	"github.com/affanhamid/go-ecommerce/utils"
)

func (c *Controller) AddAddressHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var address models.Address
	if !utils.ValidateAndDecode(w, r, &address) {
		return
	}

	if err := database.AddAddress(&address, c.DB); err != nil {

		utils.SendJSONMessage(w, http.StatusInternalServerError, "Failed to add address: "+err.Error())
		return
	}
	utils.SendJSONMessage(w, http.StatusOK, "User data added successfully!")
}

func (c *Controller) DeleteAddressHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	var addressID models.AddressID
	if !utils.ValidateAndDecode(w, r, &addressID) {
		return
	}

	if err := database.DeleteAddress(&addressID, c.DB); err != nil {

		utils.SendJSONMessage(w, http.StatusInternalServerError, "Failed to delete address: "+err.Error())
		return
	}
	utils.SendJSONMessage(w, http.StatusOK, "Address Deleted Succesfully!")
}
