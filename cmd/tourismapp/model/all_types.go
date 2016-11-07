package model

import (
	"time"
)

//import

// Clase que almacenara los datos obtenidos de cada usuario.
type Customer struct {
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	S_surname string `json:"s_surname" binding:"required"`
	Rut       string `json:"rut" binding:"required"`
	Mail      string `json:"mail" binding:"required"`
	Password  string `json:"pass" binding:"required"`
}

// Clase que almacenara los datos obtenidos de cada lugar.
type Place struct {
	Id          int64     `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Score       int64     `json:"score"`
	User_c      string    `json:"user" binding:"required"`
	Date_c      time.Time `json:"date" binding:"required"`
	Descripcion string    `json:"note" binding:"required"`
	Fono        int64     `json:"fono" binding:"required"`
}

// Clase que almacenara los datos obtenidos de cada tag.
type Tag struct {
	Id   int64  `json:"Id" binding:"required"`
	Name string `json:"Name" binding:"required"`
}

// Clase para almacenar tags de cada lugar
type Tagplace struct {
	Id_tags  int64 `json:"id_tag" binding:"required"`
	Id_place int64 `json:"id_place" binding:"required"`
}

//Horarios, 7 días de la semana, open-Close por cada día
type Schedule struct {
	Id int64     `json:"id" binding:"required"`
	O1 time.Time `json:"o1" binding:"required"`
	C1 time.Time `json:"c1" binding:"required"`
	O2 time.Time `json:"o2" binding:"required"`
	C2 time.Time `json:"c2" binding:"required"`
	O3 time.Time `json:"o3" binding:"required"`
	C3 time.Time `json:"c3" binding:"required"`
	O4 time.Time `json:"o4" binding:"required"`
	C4 time.Time `json:"c4" binding:"required"`
	O5 time.Time `json:"o5" binding:"required"`
	C5 time.Time `json:"c5" binding:"required"`
	O6 time.Time `json:"o6" binding:"required"`
	C6 time.Time `json:"c6" binding:"required"`
	O7 time.Time `json:"o7" binding:"required"`
	C7 time.Time `json:"c7" binding:"required"`
}

// Clase que permite almacenar evaluaciones.
type Evaluation struct {
	Id       int64     `json:"id" binding:"required"`
	Id_user  string    `json:"id_user" binding:"required"`
	Id_place string    `json:"id_place" binding:"required"`
	Score    int64     `json:"score" binding:"required"`
	Comment  string    `json:"comment" binding:"required"`
	Date     time.Time `json:"date" binding:"required"`
}

// Clase que almacena coordenadas y id de un lugar.
type Geocoord struct {
	Id_place int64
	Lat      float64
	Lng      float64
}
