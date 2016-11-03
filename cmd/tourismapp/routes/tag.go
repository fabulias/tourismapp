package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"tourismapp/cmd/tourismapp/model"
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
func PostTag(c *gin.Context) {
	var tag model.Tag

	//JSON enviado es enlazado a Variable del tipo Place
	err := c.Bind(&tag)
	if err != nil {
		log.Fatalln(err)
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "Missing some field required",
		}
		c.JSON(http.StatusBadRequest, response)
	} else {
		status := model.InsertTag(tag)
		if status {
			response := gin.H{
				"status":  "success",
				"data":    nil,
				"message": "Success insert",
			}
			c.JSON(http.StatusOK, response)
		} else {
			response := gin.H{
				"status":  "success",
				"data":    nil,
				"message": "Rut already exist",
			}
			c.JSON(http.StatusNotFound, response)
		}
	}
}
