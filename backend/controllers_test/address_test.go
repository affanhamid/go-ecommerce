package controllers_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/affanhamid/go-ecommerce/controllers"
	"github.com/affanhamid/go-ecommerce/database"
	"github.com/affanhamid/go-ecommerce/utils"
)

func TestAddAddress__InvalidContentType(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/api/add-user", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, c.CreateUserHandler, http.StatusUnsupportedMediaType, "Content-Type must be 'application/json'")
}

func TestAddAddress__Valid(t *testing.T) {
	db := database.Connect()
	c := controllers.Controller{DB: db}

	user, err := database.GetUser("test@example.com", db)
	if err != nil {
		log.Fatalf("Error fetching user id: " + err.Error())
	}

	address := fmt.Sprintf(`{
		"user_id": %v,
		"country": "United Kingdom",
		"state": "Greater London",
		"city": "London",
		"post_code": "W2",
		"address_line": "Fraser Place",
		"building": "383"
		}`, user.ID)

	req, err := http.NewRequest(http.MethodPost, "/api/add-address", bytes.NewBuffer([]byte(address)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	utils.Test(t, req, c.AddAddressHandler, http.StatusOK, "{\"message\":\"User data added successfully!\"}")
}

func TestGetAddresses__Valid(t *testing.T) {
	db := database.Connect()
	// c := controllers.Controller{DB: db}

	user, err := database.GetUser("test@example.com", db)
	if err != nil {
		log.Fatalf("Error fetching user id: " + err.Error())
	}

	addresses, err := database.GetAddressesFromUser(user.ID, db)
	if err != nil {
		log.Fatalf("Error fetching addresses: " + err.Error())
	}

	if len(addresses) == 0 {
		log.Fatalf("Empty address array")
	}
}

func TestDeleteAddress__Valid(t *testing.T) {
	db := database.Connect()
	c := controllers.Controller{DB: db}

	user, err := database.GetUser("test@example.com", db)
	if err != nil {
		log.Fatalf("Error fetching user id: " + err.Error())
	}

	addresses, err := database.GetAddressesFromUser(user.ID, db)
	if err != nil {
		log.Fatalf("Error fetching addresses: " + err.Error())
	}

	deleteAddress := fmt.Sprintf(`{"ID": %v}`, addresses[0].ID)

	req, err := http.NewRequest(http.MethodPost, "/api/delete-address", bytes.NewBuffer([]byte(deleteAddress)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	utils.Test(t, req, c.DeleteAddressHandler, http.StatusOK, "{\"message\":\"Address Deleted Succesfully!\"}")
}
