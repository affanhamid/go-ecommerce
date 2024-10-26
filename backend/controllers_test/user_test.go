package controllers_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/affanhamid/go-ecommerce/controllers"
	"github.com/affanhamid/go-ecommerce/database"
	"github.com/affanhamid/go-ecommerce/utils"
)

// Testing Invalid Content Type
func TestAddUser__InvalidContentType(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/api/add-user", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, c.CreateUserHandler, http.StatusUnsupportedMediaType, "Content-Type must be 'application/json'")
}

// Testing Body

func TestAddUser__InvalidJSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/api/add-user", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, c.CreateUserHandler, http.StatusBadRequest, "")
}

type TestCase struct {
	name           string
	user           string
	expectedStatus int
}

func TestAddUser__ValidJson(t *testing.T) {

	testJson := `{
		"email": "test@example.com",
		"country_code": "+1",
		"phone_number": "1234567890",
		"first_name": "John",
		"last_name": "Doe",
		"password": "password123",
		"user_type": "USER"
		}`

	req, err := http.NewRequest(http.MethodPost, "/api/add-user", bytes.NewBuffer([]byte(testJson)))
	if err != nil {
		t.Fatalf("(ValidJson) Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, c.CreateUserHandler, http.StatusOK, "{\"message\":\"User data added successfully!\"}")
}

func TestAddUser__ValidationFalseCases(t *testing.T) {
	testCases := []TestCase{
		{
			name: "Email no domain extension",
			user: `{
				"email": "test@example",
				"country_code": "+1",
				"phone_number": "1234567890",
				"first_name": "John",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Email no domain",
			user: `{
				"email": "test@.com",
				"country_code": "+1",
				"phone_number": "1234567890",
				"first_name": "John",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Email no user",
			user: `{
				"email": "@example.com",
				"country_code": "+1",
				"phone_number": "1234567890",
				"first_name": "John",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Email double .",
			user: `{
				"email": "testing@example..com",
				"country_code": "+1",
				"phone_number": "1234567890",
				"first_name": "John",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Email double @",
			user: `{
				"email": "testing@@example..com",
				"country_code": "+1",
				"phone_number": "1234567890",
				"first_name": "John",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Phone Too Big",
			user: `{
				"email": "testing@example.com",
				"country_code": "+1",
				"phone_number": "1234567890111",
				"first_name": "John",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Phone Too Small",
			user: `{
				"email": "testing@example.com",
				"country_code": "+1",
				"phone_number": "12345611",
				"first_name": "John",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Phone Non-Numeric",
			user: `{
				"email": "testing@example.com",
				"country_code": "+1",
				"phone_number": "12345676a90",
				"first_name": "John",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "First Name Too Small",
			user: `{
				"email": "testing@example.com",
				"country_code": "+1",
				"phone_number": "12345676a90",
				"first_name": "a",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "First Name Too Big",
			user: `{
				"email": "testing@example.com",
				"country_code": "+1",
				"phone_number": "12345676a90",
				"first_name": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				"last_name": "Doe",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Last Name Too Small",
			user: `{
				"email": "testing@example.com",
				"country_code": "+1",
				"phone_number": "12345676a90",
				"first_name": "John",
				"last_name": "a",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Last Name Too Big",
			user: `{
				"email": "testing@example.com",
				"country_code": "+1",
				"phone_number": "12345676a90",
				"first_name": "John",
				"last_name": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				"password": "password123",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "no password",
			user: `{
				"email": "testing@example.com",
				"country_code": "+1",
				"phone_number": "12345676a90",
				"first_name": "John",
				"last_name": "Doe",
				"password": "",
				"user_type": "USER"
				}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(http.MethodPost, "/api/add-user", bytes.NewBuffer([]byte(tc.user)))
		if err != nil {
			t.Fatalf("(%s) Could not create request: %v", tc.name, err)
		}
		req.Header.Set("Content-Type", "application/json")

		db := database.Connect()
		c := controllers.Controller{DB: db}

		utils.Test(t, req, c.CreateUserHandler, tc.expectedStatus, "")
	}
}

func TestDeleteUser__ValidJson(t *testing.T) {

	testJson := `{
		"email": "test@example.com"
		}`

	req, err := http.NewRequest(http.MethodPost, "/api/delete-user", bytes.NewBuffer([]byte(testJson)))
	if err != nil {
		t.Fatalf("(ValidJson) Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, c.DeleteUserHandler, http.StatusOK, "{\"message\":\"User deleted succesfully!\"}")
}

func TestDeleteUserPerma__ValidJson(t *testing.T) {

	testJson := `{
		"email": "test@example.com"
		}`

	req, err := http.NewRequest(http.MethodPost, "/api/delete-user-permanent", bytes.NewBuffer([]byte(testJson)))
	if err != nil {
		t.Fatalf("(ValidJson) Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, c.DeleteUserPermaHandler, http.StatusOK, "{\"message\":\"User deleted permanently succesfully!\"}")
}

func TestDeleteUser__UserDoesNotExist(t *testing.T) {

	testJson := `{
		"email": "doesNotExist@gmail.com"
		}`

	req, err := http.NewRequest(http.MethodPost, "/api/delete-user", bytes.NewBuffer([]byte(testJson)))
	if err != nil {
		t.Fatalf("(ValidJson) Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, c.DeleteUserHandler, http.StatusBadRequest, "")
}
