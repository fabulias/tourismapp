package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)













var db *sql.DB = nil
var err error

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func connectDatabase() {

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL")) //os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Error opening database: %q", err)
		DATABASE_URL := "postgres://eozcyemimcuhgg:3ac2YMMZ0EMofFw6rdrTXIky6W@ec2-107-22-250-212.compute-1.amazonaws.com:5432/da6rnltctu258a"
		db, err = sql.Open("postgres", DATABASE_URL)
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


func GetTags(c *gin.Context) {
	connectDatabase()
	pingDatabase()

	Tags := new(tags)
	rows, errq := db.Query("SELECT * FROM tags")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Tags.Id, &Tags.Name)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, Tags)
	}
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	disconnectDatabase()

}

func GetTag(c *gin.Context) {
	connectDatabase()
	pingDatabase()
	id := c.Params.ByName("id")
	Tag := new(tags)
	errq := db.QueryRow("SELECT * FROM tags WHERE id=$1", id).Scan(&Tag.Id, &Tag.Name)
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	c.JSON(http.StatusOK, Tag) //retornando los datos de user
	disconnectDatabase()
}
func GetEvaluations(c *gin.Context) {
	connectDatabase()
	pingDatabase()

	Evaluations := new(evaluation)
	rows, errq := db.Query("SELECT * FROM evaluation")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Evaluations.Id, &Evaluations.Id_user, &Evaluations.Id_place, &Evaluations.Score, &Evaluations.Comment, &Evaluations.Date)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, Evaluations)
	}
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	disconnectDatabase()

}

func GetEvaluation(c *gin.Context) {
	connectDatabase()
	pingDatabase()
	id := c.Params.ByName("id")
	Evaluations := new(evaluation)
	errq := db.QueryRow("SELECT * FROM evaluation WHERE id=$1", id).Scan(&Evaluations.Id, &Evaluations.Id_user, &Evaluations.Id_place, &Evaluations.Score, &Evaluations.Comment, &Evaluations.Date)
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	c.JSON(http.StatusOK, Evaluations) //retornando los datos de user
	disconnectDatabase()
}

func GetSchedules(c *gin.Context) { //Geo=Geolocalización
	connectDatabase()
	pingDatabase()

	Schedules := new(schedules)
	rows, errq := db.Query("SELECT * FROM schedule")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Schedules.Id, &Schedules.o1, &Schedules.c1, &Schedules.o2, &Schedules.c2, &Schedules.o3, &Schedules.c3, &Schedules.o4, &Schedules.c4, &Schedules.o5, &Schedules.c5, &Schedules.o6, &Schedules.c6, &Schedules.o7, &Schedules.c7)
		fmt.Printf("-->>%d -> %d ->%d ", Schedules.Id, Schedules.o1, Schedules.c1)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, Schedules)
	}
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	disconnectDatabase()

}

func GetSchedule(c *gin.Context) {
	connectDatabase()
	pingDatabase()
	id := c.Params.ByName("id")
	Schedules := new(schedules)
	errq := db.QueryRow("SELECT * FROM schedule WHERE id=$1", id).Scan(&Schedules.Id, &Schedules.o1, &Schedules.c1, &Schedules.o2, &Schedules.c2, &Schedules.o3, &Schedules.c3, &Schedules.o4, &Schedules.c4, &Schedules.o5, &Schedules.c5, &Schedules.o6, &Schedules.c6, &Schedules.o7, &Schedules.c7)
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	c.JSON(http.StatusOK, Schedules) //retornando los datos de user
	disconnectDatabase()
}

func GetTags_places(c *gin.Context) {
	connectDatabase()
	pingDatabase()

	Tags_places := new(tag_places)
	rows, errq := db.Query("SELECT * FROM tags_places")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Tags_places.Id_tags, &Tags_places.Id_place)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, Tags_places)
	}
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	disconnectDatabase()

}

func GetTag_place(c *gin.Context) {
	connectDatabase()
	pingDatabase()

	id := c.Params.ByName("id")
	Tags_places := new(tag_places)
	errq := db.QueryRow("SELECT * FROM tags_places WHERE id_tag=$1", id).Scan(&Tags_places.Id_tags, &Tags_places.Id_place)
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	c.JSON(http.StatusOK, Tags_places) //retornando los datos de user
	disconnectDatabase()
}

/* Función que entrega los nombres de los tag's que tiene un place*/

func TagsByPlace(c *gin.Context) {
	connectDatabase()
	pingDatabase()
	id := c.Params.ByName("id")
	Tags := new(tags)
	rows, errq := db.Query("SELECT tags.* FROM tags_places, tags WHERE tags_places.id_place=$1 AND tags.id=tags_places.id_tag ", id)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&Tags.Id, &Tags.Name)
		fmt.Printf("---->ID: %d ---->NAME: %d", Tags.Id, Tags.Name)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, Tags)
	}
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}
	disconnectDatabase()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set for next event")
		port = "8080"
	}

	r := gin.Default()

	r.Use(Cors())

	//Access localhost:port/v1
	v1 := r.Group("v1")
	{
		//Access localhost:port/v1/users
		v1.GET("/users", GetUsers)
		v1.GET("/users/:name", GetUser)
		v1.GET("/places", GetPlaces)
		v1.GET("/places/:id", GetPlace)
		v1.GET("/tags", GetTags)
		v1.GET("/tags/:id", GetTag)
		v1.GET("/evaluations", GetEvaluations)
		v1.GET("/evaluations/:id", GetEvaluation)
		v1.GET("/schedules/", GetSchedules)
		v1.GET("/schedules/:id", GetSchedule)
		v1.GET("/tags_places", GetTags_places)
		v1.GET("/tags_places/:id", GetTag_place)
		v1.GET("/tagsbyplaces/:id", TagsByPlace)
	}

	if errGin := r.Run(":" + port); errGin != nil {
		log.Printf("error listening on port "+port+": %v", errGin)
	}
}
