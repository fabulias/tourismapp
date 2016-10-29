package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tourismapp/cmd/tourismapp/model"

	_ "github.com/lib/pq"
)

//Método que busca todos los tags por lugares de la bdd.
func GetTagsPlaces(c *gin.Context) {
	tagsplaces := model.QueryTagsPlaces()
	if len(tagsplaces) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There are no tags per places",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    tagsplaces,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}
