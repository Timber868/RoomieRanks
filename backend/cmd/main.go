package main

import (
	"database/sql"
	"log"

	"github.com/Timber868/roomieranks/cmd/api"
	"github.com/Timber868/roomieranks/config"
	"github.com/Timber868/roomieranks/db"
	"github.com/go-sql-driver/mysql"
)

// Server info
const serverAddress string = ":8080"

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp", //Network protocol, tcp is standard for IP or hostname connections
		AllowNativePasswords: true,  //Allows the use of MySQL’s native password authentication method. This is often needed with more recent versions of MySQL.
		ParseTime:            true,  //Date columns wont be left as string but as Time.time instead
	})

	if err != nil {
		log.Fatal(err)
	}

	//Inititalize the database with our configuration
	initStorage(db)

	//Initialize the api server without our own custom method
	server := api.NewApiServer(serverAddress, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// Function to actually initialize the database the rest does nothing
func initStorage(db *sql.DB) {
	//Starts the database connection and checks if everything is alright
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
