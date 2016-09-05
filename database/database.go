package database

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"
)

const (
	// DBConnectionString : mysql db connection string
	DBConnectionString = "pi1:gavin@/gatorloop?charset=utf8"
)

// DB : accessible mysql database variable
var DB *sql.DB

// InitDB : initializes a database connection and variable
func InitDB() {
	if DB != nil {
		log.Error("Database has already been initialized")
		return
	}
	// connect to maria db
	var err error
	DB, err = sql.Open("mysql", DBConnectionString)
	if err != nil {
		log.Fatalf("Error opening database. %v", err)
	}

	// verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to database. %v", err)
	}
	log.Info("Connected to database.")
}
