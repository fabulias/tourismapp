package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tourismapp/cmd/model"

	_ "github.com/lib/pq"
)

//Método que busca todos los lugares de la bdd.
func GetGeocoords(c *gin.Context) {
	geocoords := model.QueryGeocoords()
	if len(geocoords) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There are no geocoords",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    geocoords,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}
