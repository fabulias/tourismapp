package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tourismapp/cmd/model"

	_ "github.com/lib/pq"
)

//Método que busca todos los lugares de la bdd.
func GetPlaces(c *gin.Context) {
	places := model.QueryPlaces()
	if len(places) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There are no places",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    places,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}
