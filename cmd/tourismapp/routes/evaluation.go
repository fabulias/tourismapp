package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tourismapp/cmd/model"

	_ "github.com/lib/pq"
)

//Método que busca todos las evaluaciones de la bdd.
func GetEvaluations(c *gin.Context) {
	evaluations := model.QueryEvaluations()
	if len(evaluations) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There are no evaluations",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    evaluations,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}
