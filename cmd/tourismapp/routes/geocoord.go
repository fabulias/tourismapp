package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tourismapp/cmd/tourismapp/model"

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

//Método que busca los datos de un usuario por su id.
func GetGeocoord(c *gin.Context) {
	id := c.Param("id")
	geocoord := model.QueryGeocoord(id)
	if len(geocoord) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no geocoord with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    geocoord,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}
