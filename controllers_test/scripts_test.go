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
	app := controllers.App{DB: db}

	utils.Test(t, req, middleware.CheckMethod("POST", app.CreateTablesHandler), http.StatusNotFound, "method not expected")
}

func TestDroptable__InvalidMethod(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/admin/dropAll", bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	app := controllers.App{DB: db}

	utils.Test(t, req, middleware.CheckMethod("POST", app.DropTablesHandler), http.StatusNotFound, "method not expected")
}

func TestDroptable__Valid(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/api/admin/dropAll", bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	app := controllers.App{DB: db}

	utils.Test(t, req, middleware.CheckMethod("POST", app.DropTablesHandler), http.StatusOK, "")
}

func TestCreatetable__Valid(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/api/admin/create", bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	db := database.Connect()
	app := controllers.App{DB: db}

	utils.Test(t, req, middleware.CheckMethod("POST", app.CreateTablesHandler), http.StatusOK, "")
}
