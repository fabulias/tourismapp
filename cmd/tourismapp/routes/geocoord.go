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

func GetGeocoordRadio(c *gin.Context) {
	id := c.Param("id")
	radio := c.Param("radio")
	place := model.QueryGeocoordRadio(id, radio)
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

func PostGeocoord(c *gin.Context) {
	var geocoord model.Geocoord

	err := c.Bind(&geocoord)
	if err != nil {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "Missing some field required",
		}
		c.JSON(http.StatusBadRequest, response)
	} else {
		status := model.InsertGeocoord(geocoord)
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
				"message": "geocoord already exist",
			}
			c.JSON(http.StatusNotFound, response)
		}
	}
}

func PatchGeocoord(c *gin.Context) {
	id := c.Param("id")
	var geo model.Geocoord
	geocoord := model.QueryGeocoord(id)

	if len(geocoord) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no geocoord with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		c.BindJSON(&geo)
		status := model.UpdateGeocoord(geo)
		if status {
			response := gin.H{
				"status":  "success",
				"data":    geo,
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
