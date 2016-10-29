package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tourismapp/cmd/tourismapp/model"

	_ "github.com/lib/pq"
)

//Método que busca todos los horarios de la bdd.
func GetSchedules(c *gin.Context) {
	schedules := model.QuerySchedules()
	if len(schedules) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There are no schedules",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    schedules,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}
