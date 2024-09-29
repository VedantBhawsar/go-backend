package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"user-crud/controllers"
)

var logger = logrus.New()
var users []controllers.User // Use the User type from the controller package

func main() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	router := mux.NewRouter()
	router.Use(loggingMiddleware)

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
