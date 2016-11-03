package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"tourismapp/cmd/tourismapp/model"
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

func PostEvaluation(c *gin.Context) {
	var evaluation model.Evaluation

	//JSON enviado es enlazado a Variable del tipo Place
	err := c.Bind(&evaluation)
	if err != nil {

		log.Fatalln(err)
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "Missing some field required",
		}
		c.JSON(http.StatusBadRequest, response)
	} else {

		status := model.InsertEvaluation(evaluation)
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
