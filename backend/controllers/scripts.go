package controllers

import (
	"net/http"

	"github.com/affanhamid/go-ecommerce/database"
)

func (c *Controller) CreateTablesHandler(w http.ResponseWriter, r *http.Request) {
	if err := database.CreateTables(c.DB); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	}
}
func (c *Controller) DropTablesHandler(w http.ResponseWriter, r *http.Request) {
	if err := database.DropTables(c.DB); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	}
}
