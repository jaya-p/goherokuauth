package goherokuauth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type authPostRequest struct {
	Username     string
	PasswordHash string
}

// authGetHandler handles GET method
func authGetHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		log.Print("token is empty")

		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "")

		return
	}

	ok, err := CheckToken(token)
	if !ok || err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "")

		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}

// authPostHandler handles POST method
func authPostHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var apr authPostRequest
	errD := decoder.Decode(&apr)
	if errD != nil {
		log.Print(errD)

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Internal Server Error\n")

		return
	}

	token, errT := GetToken(apr.Username, apr.PasswordHash)
	if token == "" || errT != nil {
		log.Print(errT)

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Internal Server Error\n")

		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, token+"\n")
}

// statusNotFoundHandler handles unknown request method
func statusNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Your requested method ("+r.Method+") is not found")
}

// authRestAPIHandler handles request response
func authRestAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Set HTTP Response "Content-Type" header as JSON
	w.Header().Set("Content-type", "text/plain; charset=utf-8")

	// Set HTTP Response Body, based on HTTP request method
	switch r.Method {
	case "GET":
		authGetHandler(w, r)
	case "POST":
		authPostHandler(w, r)
	default:
		statusNotFoundHandler(w, r)
	}
}

// RestAPIWebserver provides connectivity for handling REST API request
func RestAPIWebserver(pNum uint) {
	http.HandleFunc("/api/v1/auth", authRestAPIHandler)

	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(pNum), nil))
}
