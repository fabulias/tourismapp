package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tourismapp/cmd/model"

	_ "github.com/lib/pq"
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

/*
	Método que busca los datos de un usuario por su id.
*/
func GetUser(c *gin.Context) {
	//Verificar conexión base de datos
	//model.ConnectDatabase(Db)
	//Se obtienen parametros de la URI
	//name := c.Params.ByName("name")
	//user := new(user)
	//Se vuelve a verificar conexión para realizar consulta
	//model.PingDatabase()
	//errq := Db.QueryRow("SELECT * FROM customer WHERE name=$1", name).Scan(
	//	&user.Name,
	//	&user.Surname,
	//	&user.S_surname,
	//	&user.Rut,
	//	&user.Mail,
	//	&user.Password)
	//model.DisconnectDatabase(Db)
	//if errq != nil {
	//	log.Fatalln("Error in query ", errq)
	//Se retorna error de servidor y un mensaje
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": "There are no users",
	//	})
	//}
	c.JSON(http.StatusOK, gin.H{}) //user)
}
