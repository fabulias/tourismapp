package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"tourismapp/cmd/lib"

	_ "github.com/lib/pq"
)

/**
Clase que almacenara los datos obtenidos de cada lugar.
*/
type places struct {
	Id          int64
	Name        string
	Score       int64
	User_c      string
	Date_c      time.Time
	Descripcion string
}

/**
Método que busca todos los lugares de la bdd.
*/
func GetPlaces(c *gin.Context) {
	//Verificar conexión base de datos
	lib.ConnectDatabase()
	//Objeto a recibir cada fila de la consulta
	Place := new(places)
	//Se vuelve a verificar conexión para realizar consulta
	lib.PingDatabase()
	//Consulta a base de datos
	rows, errq := db.Query("SELECT * FROM place")
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
		err := rows.Scan(&Place.Id,
			&Place.Name,
			&Place.Score,
			&Place.User_c,
			&Place.Date_c,
			&Place.Descripcion)
		if err != nil {
			log.Fatal(err)
		}
		//Respuesta del servidor
		c.JSON(http.StatusOK, Place)
	}
}

/**
Método que busca los datos de un lugar por su id.
*/
func GetPlace(c *gin.Context) {
	//Verificar conexión base de datos
	lib.ConnectDatabase()
	//Se obtienen parametros de la URI
	id := c.Params.ByName("id")
	Place := new(places)
	//Se vuelve a verificar conexión para realizar consulta
	lib.PingDatabase()
	errq := db.QueryRow("SELECT * FROM place WHERE id=$1", id).Scan(
		&Place.Id,
		&Place.Name,
		&Place.Score,
		&Place.User_c,
		&Place.Date_c,
		&Place.Descripcion)
	lib.DisconnectDatabase()
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		//Se retorna error de servidor y un mensaje
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are no users",
		})
	}
	c.JSON(http.StatusOK, Place)
}
