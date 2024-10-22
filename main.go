package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	routes.LoadRoutes()

	fmt.Printf("Starting server on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
