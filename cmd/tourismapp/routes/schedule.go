package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"tourismapp/cmd/tourismapp/model"
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
func PostSchedule(c *gin.Context) {
	var schedule model.Schedule

	//JSON enviado es enlazado a Variable del tipo Place
	err := c.Bind(&schedule)
	if err != nil {

		log.Fatalln(err)
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "Missing some field required",
		}
		c.JSON(http.StatusBadRequest, response)
	} else {
		status := model.InsertSchedule(schedule)
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
				"message": "Place's Schedule already exist",
			}
			c.JSON(http.StatusNotFound, response)
		}
	}

}
