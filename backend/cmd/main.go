package main

import (
	"database/sql"
	"log"

	"github.com/Timber868/roomieranks/cmd/api"
	"github.com/Timber868/roomieranks/config"
	"github.com/Timber868/roomieranks/db"
	"github.com/go-sql-driver/mysql"
)

const serverAddress string = ":8080"

func main() {
	// Initialize MySQL storage connection
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp", // Network protocol
		AllowNativePasswords: true,  // Allows MySQL native password authentication
		ParseTime:            true,  // Date columns are parsed as time
	})

	if err != nil {
		log.Fatal(err)
	}

	// Initialize the storage (i.e., database connection)
	initStorage(db)

	// Create a new API server instance
	server := api.NewApiServer(serverAddress, db)

	// Start the server with the CORS-enabled routes
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	// Verify the database connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
