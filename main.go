package main

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func connectDatabase() {
	var err error
	db, err = sql.Open("postgres", "postgres://eozcyemimcuhgg:3ac2YMMZ0EMofFw6rdrTXIky6W@ec2-107-22-250-212.compute-1.amazonaws.com:5432/da6rnltctu258a")//os.Getenv("DATABASE_URL"))

	if err != nil {
  	log.Fatalln("Error opening database: %q", err)
  }
}

func disconnectDatabase() {
	err := db.Close()
	if err != nil {
		log.Fatalln("Error closing database: %q", err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetUsers(c *gin.Context)  {
	c.JSON(200, gin.H{"Ping":"Pong"})
}

func main() {

	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("v1")
	{
		v1.GET("/users", GetUsers)
		//v1.GET("/users/:id", GetUser)
	}

	ret := db.Ping()
	if ret != nil {
		//
	}
	r.Run(":8080")
}
