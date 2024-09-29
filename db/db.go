package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDb(connStr string) {
	var err error

	// Initialize the database connection
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open the database:", err)
	}

	// Test the connection
	err = Db.Ping()
	if err != nil {
		log.Fatal("Unable to connect to the database:", err)
	}

	log.Println("Connected to the database successfully!")
}
