package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-crud/controllers"
	"user-crud/db"
	"user-crud/middlewares"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var users []controllers.User // Use the User type from the controller package

func main() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	router := mux.NewRouter()
	router.Use(middlewares.LoggingMiddleware)
	db.ConnectDb("postgres://postgres:password@localhost:5432/postgres?sslmode=disable")

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to the Go backend API")
	})

	fmt.Println("Server started on port 8080")
	logger.Info("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
