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

	v1 := r.Group("api/v1")
	{
		v1.GET("/customers", routes.GetUsers)
		v1.GET("/customers/:rut", routes.GetUser)
		v1.GET("/evaluations", routes.GetEvaluations)
		v1.GET("/evaluations/:id", routes.GetEvaluation)
		v1.GET("/geocoords", routes.GetGeocoords)
		v1.GET("/geocoords/:id", routes.GetGeocoord)
		v1.GET("/places", routes.GetPlaces)
		v1.GET("/places/:id", routes.GetPlace)
		v1.GET("/schedules", routes.GetSchedules)
		v1.GET("/schedules/:id", routes.GetSchedule)
		v1.GET("/tags", routes.GetTags)
		v1.GET("/tags/:id", routes.GetTag)
		v1.GET("/tags_places", routes.GetTagsPlaces)
		v1.GET("/tags_place", routes.GetTagPlace)

		v1.POST("/customers", routes.PostUser)
		v1.POST("/evaluations", routes.PostEvaluation)
		v1.POST("/geocoords", routes.PostGeocoord)
		v1.POST("/places", routes.PostPlace)
		v1.POST("/schedules", routes.PostSchedule)
		v1.POST("/tags", routes.PostTag)
		v1.POST("/tags_places", routes.PostTagPlace)

		v1.PATCH("/customers/:rut", routes.PatchUser)
		v1.PATCH("/evaluations/:id", routes.PatchEvaluation)
		v1.PATCH("/geocoords/:id", routes.PatchGeocoord)
		v1.PATCH("/places/:id", routes.PatchPlace)
		v1.PATCH("/schedules/:id", routes.PatchSchedule)
		v1.PATCH("/tags/:id", routes.PatchTag)
		/*


			v1.DELETE("/customers", routes.)
			v1.DELETE("/customers/:rut", routes.)
			v1.DELETE("/evaluations", routes.)
			v1.DELETE("/evaluations/:id", routes.)
			v1.DELETE("/geocoords", routes.)
			v1.DELETE("/geocoords/:id", routes.)
			v1.DELETE("/places", routes.)
			v1.DELETE("/places/:id", routes.)
			v1.DELETE("/schedules", routes.)
			v1.DELETE("/schedules/:id", routes.)
			v1.DELETE("/tags", routes.)
			v1.DELETE("/tags/:id", routes.)
			v1.DELETE("/tags_places", routes.)
			v1.DELETE("/tags_place", routes.)
		*/
	}
	log.Println("Uploading...", time.Now())
	log.Println("Running on port : " + port)
	if errGin := r.Run(":" + port); errGin != nil {
		log.Printf("error listening on port "+port+": %v", errGin)
	}
}
