package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var logger = logrus.New()

var users []User

func main() {
	// Initialize users slice
	users = []User{}
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/create", createUser).Methods("POST")
	router.HandleFunc("/{id}", getUser).Methods("GET")
	router.HandleFunc("/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to the Go backend API")
	})
	fmt.Println("Server started on port 8080")
	// Start the server
	loggedMux := loggingMiddleware(router)

	logrus.Info("Starting server on :8080")
	http.ListenAndServe(":8080", loggedMux)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.Path,
			"time":   time.Since(start),
		}).Info("handled request")
	})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getUsers")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createUser")
	// Logic to create a user should be added here
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getUser")
	// Logic to retrieve a specific user by ID should be added here
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUser")
	// Logic to update a user should be added here
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteUser")
	// Logic to delete a user should be added here
	json.NewEncoder(w).Encode(users)
}
