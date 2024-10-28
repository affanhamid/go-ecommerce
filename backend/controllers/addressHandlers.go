package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (c *Controller) GetAddressesHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	query := r.URL.Query()

	userId := query.Get("userId")

	if userId == "" {
		utils.SendJSONMessage(w, http.StatusBadRequest, "no user id provided")
	}

	num, err := strconv.ParseUint(userId, 10, 0)
	if err != nil {
		utils.SendJSONMessage(w, http.StatusInternalServerError, "cannot convert user id to uint")
		return
	}

	addresses, err := database.GetAddressesFromUser(uint(num), c.DB)
	if err != nil {

		utils.SendJSONMessage(w, http.StatusInternalServerError, "Failed to retrieve addresses: "+err.Error())
		return
	}

	jsonBytes, err := json.Marshal(addresses)
	if err != nil {
		utils.SendJSONMessage(w, http.StatusInternalServerError, "Failed to marshal json: "+err.Error())
	}
	utils.SendJSON(w, http.StatusOK, jsonBytes)
}

func (c *Controller) DeleteAddressHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateJSONContentType(w, r) {
		return
	}

	query := r.URL.Query()
	addressId := query.Get("address_id")

	if addressId == "" {
		utils.SendJSONMessage(w, http.StatusBadRequest, "No address id provided")
		return
	}

	num, err := strconv.ParseUint(addressId, 10, 0)
	if err != nil {
		utils.SendJSONMessage(w, http.StatusInternalServerError, "cannot convert address id to uint")
		return
	}

	if err := database.DeleteAddress(uint(num), c.DB); err != nil {

		utils.SendJSONMessage(w, http.StatusInternalServerError, "Failed to delete address: "+err.Error())
		return
	}
	utils.SendJSONMessage(w, http.StatusOK, "Address Deleted Succesfully!")
}
