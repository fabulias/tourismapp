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

//Método que busca los datos de una evaluación por su id.
func GetEvaluation(c *gin.Context) {
	id := c.Param("id")
	evaluation := model.QueryEvaluation(id)
	if len(evaluation) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no evaluation with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    evaluation,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}

func PostEvaluation(c *gin.Context) {
	var evaluation model.Evaluation
	err := c.Bind(&evaluation)

	if err != nil {
		log.Println(err)
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
				"message": "id evaluation already exist",
			}
			c.JSON(http.StatusNotFound, response)
		}
	}
}

func PatchEvaluation(c *gin.Context) {
	rut := c.Param("id")
	var eval model.Evaluation
	evaluation := model.QueryEvaluation(rut)

	if len(evaluation) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no eval with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		c.BindJSON(&eval)
		status := model.UpdateEvaluation(eval)
		if status {
			response := gin.H{
				"status":  "success",
				"data":    eval,
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

func DeleteEvaluation(c *gin.Context) {
	id := c.Param("id")
	var eval model.Evaluation
	evaluation := model.QueryEvaluation(id)

	if len(evaluation) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There is no evaluation with that id",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		c.BindJSON(&eval)
		status := model.EraseEvaluation(id)
		if status {
			response := gin.H{
				"status":  "success",
				"data":    nil,
				"message": "",
			}
			c.JSON(http.StatusOK, response)
		} else {
			response := gin.H{
				"status":  "error",
				"data":    nil,
				"message": "Erase failed!",
			}
			c.JSON(http.StatusNotFound, response)
		}
	}
}
