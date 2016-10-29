package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tourismapp/cmd/tourismapp/model"

	_ "github.com/lib/pq"
)

//Método que busca todos los tags de la bdd.
func GetTags(c *gin.Context) {
	tags := model.QueryTags()
	if len(tags) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There are no tags",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    tags,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}
