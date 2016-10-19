package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"database/sql"

	_ "github.com/lib/pq"
)

/*
	Clase
	**/

type places struct {
	Id          int64     `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Score       int64     `json:"score" binding:"required"`
	User_c      string    `json:"user_c" binding:"required"`
	Date_c      time.Time `json:"fecha_c" binding:"required"`
	Descripcion string    `json: "descripcion" binding:"required"`
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
