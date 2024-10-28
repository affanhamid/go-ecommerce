package controllers_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/affanhamid/go-ecommerce/controllers"
	"github.com/affanhamid/go-ecommerce/database"
	"github.com/affanhamid/go-ecommerce/middleware"
	"github.com/affanhamid/go-ecommerce/utils"
)

func TestCreatetable__InvalidMethod(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/admin/create", bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, middleware.CheckMethod("POST", c.CreateTablesHandler), http.StatusNotFound, "method not expected")
}

func TestDropTable__InvalidMethod(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/admin/dropAll", bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, middleware.CheckMethod("POST", c.DropTablesHandler), http.StatusNotFound, "method not expected")
}

func TestDropTable__Valid(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/api/admin/dropAll", bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, middleware.CheckMethod("POST", c.DropTablesHandler), http.StatusOK, "")
}

func TestCreatetable__Valid(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/api/admin/create", bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	c := controllers.Controller{DB: db}

	utils.Test(t, req, middleware.CheckMethod("POST", c.CreateTablesHandler), http.StatusOK, "")
}
