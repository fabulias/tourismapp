package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tourismapp/cmd/lib"

	_ "github.com/lib/pq"
)

/**
Clase que almacenara los datos obtenidos de cada usuario.
*/
type user struct {
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	S_surname string `json:"s_surname" binding:"required"`
	Rut       string `json:"rut" binding:"required"`
	Mail      string `json:"mail" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

/**
Método que busca todos los usuarios de la bdd.
*/
func GetUsers(c *gin.Context) {
	log.Println("Entro a users")
	//Verificar conexión base de datos
	lib.ConnectDatabase()
	//Objeto a recibir cada fila de la consulta
	User := new(user)
	//Se vuelve a verificar conexión para realizar consulta
	lib.PingDatabase()
	//Consulta a base de datos
	rows, errq := db.Query("SELECT * FROM customer")
	//Desconectarse de la base de datos.
	lib.DisconnectDatabase()
	// Si la consulta no funciona se retorna error de servidor
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}
	defer rows.Close()
	// Recorrer las filas retornadas
	for rows.Next() {
		// Se guarda valor en variable User
		err := rows.Scan(&User.Name,
			&User.Surname,
			&User.S_surname,
			&User.Rut,
			&User.Mail,
			&User.Password)
		if err != nil {
			log.Fatal(err)
		}
		//Respuesta del servidor
		c.JSON(http.StatusOK, User)
	}
}

/*
	Método que busca los datos de un usuario por su id.
*/
func GetUser(c *gin.Context) {
	//Verificar conexión base de datos
	lib.ConnectDatabase()
	//Se obtienen parametros de la URI
	name := c.Params.ByName("name")
	user := new(user)
	//Se vuelve a verificar conexión para realizar consulta
	lib.PingDatabase()
	errq := db.QueryRow("SELECT * FROM customer WHERE name=$1", name).Scan(
		&user.Name,
		&user.Surname,
		&user.S_surname,
		&user.Rut,
		&user.Mail,
		&user.Password)
	lib.DisconnectDatabase()
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		//Se retorna error de servidor y un mensaje
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}
	c.JSON(http.StatusOK, user)
}
