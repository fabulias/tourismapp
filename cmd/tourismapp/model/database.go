package model

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var (
	err error
	db  *sql.DB
)

func connectDatabase() {
	db, err = sql.Open("postgres", "postgres://eozcyemimcuhgg:3ac2YMMZ0EMofFw6rdrTXIky6W@ec2-107-22-250-212.compute-1.amazonaws.com:5432/da6rnltctu258a") //os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Error opening database: %q", err)
		DATABASE_URL := "postgres://eozcyemimcuhgg:3ac2YMMZ0EMofFw6rdrTXIky6W@ec2-107-22-250-212.compute-1.amazonaws.com:5432/da6rnltctu258a"
		db, err = sql.Open("postgres", DATABASE_URL)
	}
}

func disconnectDatabase() {
	err = db.Close()
	if err != nil {
		log.Println("Error closing database: %q", err)
	}
}

func pingDatabase() {
	err = db.Ping()
	if err != nil {
		log.Println("Error ping to database", err)
	}
}
