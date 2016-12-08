package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	//"time"
	"tourismapp/cmd/tourismapp/model"
)

func checkFieldSc(sc model.Schedule) bool {
	if sc.Id != 0 {
		return false
	}
	return true
}

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

//Método que busca los datos de un horario por su id.
func GetSchedule(c *gin.Context) {
	id := c.Param("id")
	schedule := model.QuerySchedule(id)
	if len(schedule) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no schedule with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    schedule,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}

func PostSchedule(c *gin.Context) {
	var schedule model.Schedule

	//JSON enviado es enlazado a Variable del tipo Schedule
	err := c.BindJSON(&schedule)
	if err != nil {
		if !checkFieldSc(schedule) {
			response := gin.H{
				"status":  "error",
				"data":    nil,
				"message": "Missing some field required",
			}
			c.JSON(http.StatusBadRequest, response)
		}
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

func PatchSchedule(c *gin.Context) {
	id := c.Param("id")
	var user model.Schedule
	schedule := model.QuerySchedule(id)

	if len(schedule) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no user with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		c.BindJSON(&user)
		status := model.UpdateSchedule(user)
		if status {
			response := gin.H{
				"status":  "success",
				"data":    user,
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
