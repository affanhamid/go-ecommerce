package controllers

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/database"
)

func (app *App) CreateTablesHandler(w http.ResponseWriter, r *http.Request) {
	if err := database.CreateTables(app.DB); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	}
}
func (app *App) DropTablesHandler(w http.ResponseWriter, r *http.Request) {
	if err := database.DropTables(app.DB); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	}
}
