package routes

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	//"log"
	"net/http"
	"tourismapp/cmd/tourismapp/model"
)

//Método que busca todos los usuarios de la bdd.
func GetUsers(c *gin.Context) {
	customers := model.QueryCustomers()
	if len(customers) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There are no users",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    customers,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}

//Método que busca los datos de un usuario por su id.
func GetUser(c *gin.Context) {
	rut := c.Param("rut")
	customer := model.QueryCustomer(rut)
	if len(customer) == 0 {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "There ir no user with that rut",
		}
		c.JSON(http.StatusNotFound, response)
	} else {
		response := gin.H{
			"status":  "success",
			"data":    customer,
			"message": nil,
		}
		c.JSON(http.StatusOK, response)
	}
}

//Método que inserta un nuevo usuario en la base de datos
func PostUser(c *gin.Context) {
	var user model.Customer

	//JSON enviado es enlazado a Variable del tipo Customer user
	err := c.Bind(&user)
	if err != nil {
		response := gin.H{
			"status":  "error",
			"data":    nil,
			"message": "Missing some field",
		}
		c.JSON(http.StatusOK, response)
	} else {
		request := model.InsertCustomer(user)
		if request == true {
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
			c.JSON(http.StatusOK, response)
		}
	}

}
