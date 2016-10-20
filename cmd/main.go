package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

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
		/*
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
		*/
	}

	if errGin := r.Run(":" + port); errGin != nil {
		log.Printf("error listening on port "+port+": %v", errGin)
	}
}
