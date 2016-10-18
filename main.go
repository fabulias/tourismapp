package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil
var err error

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func connectDatabase() {
	db, err = sql.Open("postgres", "postgres://eozcyemimcuhgg:3ac2YMMZ0EMofFw6rdrTXIky6W@ec2-107-22-250-212.compute-1.amazonaws.com:5432/da6rnltctu258a") //os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Error opening database: %q", err)
	}
}

func disconnectDatabase() {
	err = db.Close()
	if err != nil {
		log.Fatalln("Error closing database: %q", err)
	}
}

func pingDatabase() {
	err = db.Ping()
	if err != nil {
		log.Fatalln("Error ping to database", err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetUsers(c *gin.Context) {
	connectDatabase()
	pingDatabase()
	var name string
	errq := db.QueryRow("SELECT surname FROM customer WHERE name='nicolas'").Scan(&name)

	fmt.Printf("Nombre", name)

	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	//defer rows.Close()
	disconnectDatabase()
	c.JSON(http.StatusOK, gin.H{"Ping": "Pong"})
}

func GetUser(c *gin.Context) {

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set for next event")
		port = "8080"
	}

	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
	}

	if errGin := r.Run(":" + port); errGin != nil {
		log.Printf("error listening on port "+port+": %v", errGin)
	}
}
