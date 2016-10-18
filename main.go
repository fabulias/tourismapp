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

type user struct {
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	S_surname string `json:"s_surname" binding:"required"`
	Rut       string `json:"rut" binding:"required"`
	Mail      string `json:"mail" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type places struct {
	Id     int64  `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Score  int64  `json:"score" binding:"required"`
	User_c string `json:"user_c" binding:"required"`
	//Fecha_c     time.Time `json:"fecha_c" binding:"required"`
	Descripcion string `json: "descripcion" binding:"required"`
}

type evaluation struct {
	Id       int64 `json:"id" binding:"required"`
	Id_user  int64 `json:"id_user" binding:"required"`
	Id_place int64 `json:"id_place" binding:"required"`
	Score    int64 `json:"score" binding:"required"`
	//Date     time.Time `json:"date" binding:required`
}

type schedules struct { //Horarios, 7 días de la semana, open-Close por cada día
	Id int64 `json:"id_place" binding:"required"`
	/*	o1 time.Time `json:"o1" binding:"required"`
		c1 time.Time `json:"c1" binding:"required"`
		o2 time.Time `json:"o2" binding:"required"`
		c2 time.Time `json:"c2" binding:"required"`
		o3 time.Time `json:"o3" binding:"required"`
		c3 time.Time `json:"c3" binding:"required"`
		o4 time.Time `json:"o4" binding:"required"`
		c4 time.Time `json:"c4" binding:"required"`
		o5 time.Time `json:"o5" binding:"required"`
		c5 time.Time `json:"c5" binding:"required"`
		o6 time.Time `json:"o6" binding:"required"`
		c6 time.Time `json:"c6" binding:"required"`
		o7 time.Time `json:"o7" binding:"required"`
		c7 time.Time `json:"c7" binding:"required"` */
}

type geocord struct {
	Id_place int64 `json:"id_place" binding:"required"`
	latitude int64 `json:"latitud" binding :"required"`
	altitude int64 `json:"altitud" binding:"required"`
}

type tag_places struct {
	Id_tags  int64 `json:"id_tags" binding:"required"`
	Id_place int64 `json:"id_place" binding:"required"`
}
type tags struct {
	Id   int64  `json:"id_tags" binding:"required"`
	Name string `json:"name" binding:"required"`
}

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
	//var []user = new(user)
	rows, errq := db.Query("SELECT name FROM customer")
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("----> %s", name)
	}

	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	//defer rows.Close()
	disconnectDatabase()
	c.JSON(http.StatusOK, gin.H{"Ping": "hola"})
}

func GetUser(c *gin.Context) {
	connectDatabase()
	pingDatabase()

	name := c.Params.ByName("name")
	fmt.Printf("-------->%s", name)
	var user = new(user)
	errq := db.QueryRow("SELECT * FROM customer WHERE name=$1", name).Scan(&user.Name, &user.Surname, &user.S_surname, &user.Rut, &user.Mail, &user.Password)

	fmt.Printf("Nombre->%d", user.Name)

	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	//defer rows.Close()
	disconnectDatabase()
	c.JSON(http.StatusOK, gin.H{"Name": user.Name,
		"Surname":   user.Surname,
		"S_surname": user.S_surname,
		"Rut":       user.Rut,
		"Mail":      user.Mail,
		"Pass":      user.Password})
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
		v1.GET("/users/:name", GetUser)
	}

	if errGin := r.Run(":" + port); errGin != nil {
		log.Printf("error listening on port "+port+": %v", errGin)
	}
}
