package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/affanhamid/go-ecommerce/app"
	"github.com/affanhamid/go-ecommerce/controllers"
	"github.com/affanhamid/go-ecommerce/database"
	"github.com/affanhamid/go-ecommerce/routes"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv: ", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("Port not present in environment, using 8080")
		port = "8080"
	}

	db := database.Connect()
	c := controllers.Controller{DB: db}
	app := app.App{Controller: &c}

	routes.LoadAdminRoutes(&app)
	routes.LoadUserRoutes(&app)

	fmt.Printf("Starting server on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
