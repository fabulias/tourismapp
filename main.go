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

type user struct {
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	S_surname string `json:"s_surname" binding:"required"`
	Rut       string `json:"rut" binding:"required"`
	Mail      string `json:"mail" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type places struct {
	Id          int64     `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Score       int64     `json:"score" binding:"required"`
	User_c      string    `json:"user_c" binding:"required"`
	Date_c      time.Time `json:"fecha_c" binding:"required"`
	Descripcion string    `json: "descripcion" binding:"required"`
}

type evaluation struct {
	Id       int64     `json:"id" binding:"required"`
	Id_user  string    `json:"id_user" binding:"required"`
	Id_place string    `json:"id_place" binding:"required"`
	Score    int64     `json:"score" binding:"required"`
	Comment  string    `json:"comment" binding:"required"`
	Date     time.Time `json:"date" binding:required`
}

type schedules struct { //Horarios, 7 días de la semana, open-Close por cada día
	Id int64  `json:"id_place" binding:"required"`
	o1 string `json:"o1" binding:"required"`
	c1 string `json:"c1" binding:"required"`
	o2 string `json:"o2" binding:"required"`
	c2 string `json:"c2" binding:"required"`
	o3 string `json:"o3" binding:"required"`
	c3 string `json:"c3" binding:"required"`
	o4 string `json:"o4" binding:"required"`
	c4 string `json:"c4" binding:"required"`
	o5 string `json:"o5" binding:"required"`
	c5 string `json:"c5" binding:"required"`
	o6 string `json:"o6" binding:"required"`
	c6 string `json:"c6" binding:"required"`
	o7 string `json:"o7" binding:"required"`
	c7 string `json:"c7" binding:"required"`
}

type geocord struct {
	Id_place int64 `json:"id_place" binding:"required"`
	pos      int64 `json:"latitud" binding :"required"`
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
	//User := user{Name: " ", Surname: " ", S_surname: " ", Rut: " ", Mail: " ", Password: " "}
	User := new(user)                                //Creando objeto USER
	rows, errq := db.Query("SELECT * FROM customer") //Resultado de consulta se guarda en ROWS. Query retornará multiples filas de la BD
	defer rows.Close()                               //Se cierra el ROWS para posteriormente recorrerlo
	for rows.Next() {                                //For que recorre rows
		//var name string
		err := rows.Scan(&User.Name, &User.Surname, &User.S_surname, &User.Rut, &User.Mail, &User.Password) //Se guardan datos de rows en User
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, User) //Retornando los datos de user
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
}

func GetUser(c *gin.Context) {
	connectDatabase()
	pingDatabase()

	name := c.Params.ByName("name") //Obtiene el parametro enviado por url llamado name.
	user := new(user)               // Creando object user
	errq := db.QueryRow("SELECT * FROM customer WHERE name=$1", name).Scan(&user.Name, &user.Surname, &user.S_surname, &user.Rut, &user.Mail, &user.Password)
	//QueryRow devuelve sólo una fila de la DB, la cual se almacena en user
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	c.JSON(http.StatusOK, user) //retornando los datos de user
	disconnectDatabase()
}

func GetPlaces(c *gin.Context) {
	connectDatabase()
	pingDatabase()

	Place := new(places)
	rows, errq := db.Query("SELECT * FROM place")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Place.Id, &Place.Name, &Place.Score, &Place.User_c, &Place.Date_c, &Place.Descripcion)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, Place)
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

func GetPlace(c *gin.Context) {
	connectDatabase()
	pingDatabase()
	id := c.Params.ByName("id")
	Place := new(places)
	errq := db.QueryRow("SELECT * FROM place WHERE id=$1", id).Scan(&Place.Id, &Place.Name, &Place.Score, &Place.User_c, &Place.Date_c, &Place.Descripcion)
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		disconnectDatabase()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}

	c.JSON(http.StatusOK, Place) //retornando los datos de user
	disconnectDatabase()
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

	v1 := r.Group("v1")
	{
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
