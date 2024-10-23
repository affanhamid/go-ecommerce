package middleware

import (
	"net/http"
)

func CheckMethod(method string, routeHandler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "method not expected ", http.StatusNotFound)
			return
		}
		routeHandler(w, r)
	}
}
