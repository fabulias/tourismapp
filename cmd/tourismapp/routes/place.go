package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"tourismapp/cmd/tourismapp/model"
)

//Método que busca todos los lugares de la bdd.
func GetPlaces(c *gin.Context) {
	places := model.QueryPlaces()
	if len(places) == 0 {
		response := gin.H{
			"status":   "error",
			"data":     nil,
			"messageS": "There are no places",
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

//Método que busca los datos de un lugar por su id.
func GetPlace(c *gin.Context) {
	id := c.Param("id")
	place := model.QueryPlace(id)
	if len(place) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no place with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    place,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}

func PostPlace(c *gin.Context) {
	var place model.Place
	err := c.Bind(&place)
	if err != nil {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "Missing some field required",
		}
		c.JSON(http.StatusBadRequest, response)
	} else {
		status := model.InsertPlace(place)
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
				"message": "place already exist",
			}
			c.JSON(http.StatusNotFound, response)
		}
	}
}

func PatchPlace(c *gin.Context) {
	id := c.Param("id")
	var pl model.Place
	place := model.QueryPlace(id)

	if len(place) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no place with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		c.BindJSON(&pl)
		status := model.UpdatePlace(pl)
		if status {
			response := gin.H{
				"status":  "success",
				"data":    pl,
				"message": "",
			}
			c.JSON(http.StatusOK, response)
		} else {
			response := gin.H{
				"status":  "success",
				"data":    nil,
				"message": "Upload failed!",
			}
			c.JSON(http.StatusNotFound, response)
		}
	}
}
