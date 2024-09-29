package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id    int    `json:"id"`
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
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createUser")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	var user User
	var userLength int = len(users)
	user.Id = userLength + 1
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Unable to unmarshal body", http.StatusBadRequest)
		return
	}
	users = append(users, user)

	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println(id)
	for _, user := range users {
		if userId, err := strconv.Atoi(id); err == nil && user.Id == userId {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("User not found")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUser")
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	// Temporary struct with unexported fields
	type Temp struct {
		Name  *string `json:"name"`
		Age   *int    `json:"age"`
		Email *string `json:"email"`
	}

	var user1 Temp
	fmt.Printf("Request Body: %s\n", string(body))
	err = json.Unmarshal(body, &user1)
	if err != nil {
		http.Error(w, "Unable to unmarshal body", http.StatusBadRequest)
		return
	}

	for index, user := range users {
		if userId, err := strconv.Atoi(id); err == nil && user.Id == userId {
			if user1.Email != nil {
				users[index].Email = *user1.Email
			}
			if user1.Age != nil {
				users[index].Age = *user1.Age
			}
			if user1.Name != nil {
				users[index].Name = *user1.Name
			}
			json.NewEncoder(w).Encode(users[index])
			return
		}
	}

	json.NewEncoder(w).Encode("User not found")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteUser")
	vars := mux.Vars(r)
	id := vars["id"]
	var error = true
	for index, user := range users {
		if userId, err := strconv.Atoi(id); err == nil && user.Id == userId {
			users = append(users[:index], users[index+1:]...)
			error = false
		}
	}
	if error {
		json.NewEncoder(w).Encode("User not found")
		return
	}
	json.NewEncoder(w).Encode(users)
	return
}
