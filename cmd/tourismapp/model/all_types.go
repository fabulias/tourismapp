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
	Id int64
	O1 time.Time
	C1 time.Time
	O2 time.Time
	C2 time.Time
	O3 time.Time
	C3 time.Time
	O4 time.Time
	C4 time.Time
	O5 time.Time
	C5 time.Time
	O6 time.Time
	C6 time.Time
	O7 time.Time
	C7 time.Time
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
