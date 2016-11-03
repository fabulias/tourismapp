package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
	"tourismapp/cmd/tourismapp/routes"
)

var DEFAULT_PORT string = "8080"

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT must be set for next event...")
		log.Println("$PORT = DEFAULT_PORT")
		port = DEFAULT_PORT
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(Cors())
	//Access localhost:port/v1

	v1 := r.Group("v1")
	{
		//Access localhost:port/v1/users
		v1.GET("/users", routes.GetUsers)
		v1.GET("/users/:rut", routes.GetUser)
		v1.GET("/places", routes.GetPlaces)
		v1.GET("/tags", routes.GetTags)
		v1.GET("/evaluations", routes.GetEvaluations)
		v1.GET("/schedules/", routes.GetSchedules)
		v1.GET("/tags_places", routes.GetTagsPlaces)

		v1.POST("/users", routes.PostUser)
		v1.POST("/places", routes.PostPlace)
		v1.POST("/tags", routes.PostTag)
		v1.POST("/tags_places", routes.PostTagPlace)
		v1.POST("/evaluations", routes.PostEvaluation)
		v1.POST("/schedules", routes.PostSchedule)

		//v1.GET("/geocoords", routes.GetGeocoords)
		//v1.POST("/users/:rut", routes.PostUser)
		/*
			v1.GET("/places/:id", GetPlace)
			v1.GET("/tags/:id", GetTag)
			v1.GET("/evaluations/:id", GetEvaluation)
			v1.GET("/schedules/:id", GetSchedule)
			v1.GET("/tags_places/:id", GetTag_place)
			v1.GET("/tagsbyplaces/:id", TagsByPlace)
		*/
	}
	log.Println("Uploading...", time.Now())
	log.Println("Running on port : " + port)
	if errGin := r.Run(":" + port); errGin != nil {
		log.Printf("error listening on port "+port+": %v", errGin)
	}
}
