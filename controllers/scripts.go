package controllers

import (
	"fmt"
	"net/http"

	"github.com/affanhamid/go-ecommerce/database"
)

func CreateTablesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, database.CreateTables())
}

func DropTablesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, database.DropAllTables())
}
