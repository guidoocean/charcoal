package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Extend(r *mux.Router) error {
	prefixed := r.PathPrefix("/signup/").Subrouter()
	prefixed.HandleFunc("/", renderSignupForm).Methods("GET")
	prefixed.HandleFunc("/", performSignup).Methods("POST")
	return nil
}

func renderSignupForm(w http.ResponseWriter, r *http.Request) {

}

func performSignup(w http.ResponseWriter, r *http.Request) {

}
