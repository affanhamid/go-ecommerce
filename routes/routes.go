package routes

import (
	"net/http"
)

func LoadRoutes() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
}
