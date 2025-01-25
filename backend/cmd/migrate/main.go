package main

import (
	"log"
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Timber868/roomieranks/config"
	"github.com/Timber868/roomieranks/db"
	"github.com/golang-migrate/migrate/v4"
)

//This is where I handle all our migrations
//We will use the migrate package to handle our mysql database

func main() {
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
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

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	//Two types Up or down, if you want to do changes up and revert is down
	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	} else if cmd == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}
}

// Running this in up table will create the table
// CREATE TABLE IF NOT EXISTS users (
//     `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
//     `firstName` VARCHAR(255) NOT NULL,
//     `lastName` VARCHAR(255) NOT NULL,
//     `email` VARCHAR(255) NOT NULL,
//     `password` VARCHAR(255) NOT NULL,
//     `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

//     PRIMARY KEY (`id`),
//     UNIQUE KEY `email` (`email`)
// )

// For the down:
// DROP TABLE IF EXISTS users;
