package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"tourismapp/cmd/tourismapp/model"
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

//Método que busca los datos de un usuario por su id.
func GetTagPlace(c *gin.Context) {
	idp := c.Query("idp")
	idt := c.Query("idt")
	if idp == "" || idt == "" {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "I can't found idp and idt :( ",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tag := model.QueryTagsPlace(idp, idt)
	if len(tag) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no tag with those IDs",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    tag,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}

func PostTagPlace(c *gin.Context) {
	var tagplace model.Tagplace

	//JSON enviado es enlazado a Variable del tipo Place
	err := c.Bind(&tagplace)
	if err != nil {
		log.Fatalln(err)
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "Missing some field required",
		}
		c.JSON(http.StatusBadRequest, response)
	} else {
		status := model.InsertTagPlace(tagplace)
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
