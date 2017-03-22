package model

import (
	"time"
)

//import

// Clase que almacenara los datos obtenidos de cada usuario.


// Clase que almacenara los datos obtenidos de cada lugar.
type Place struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Score       int64     `json:"score"`
	User_c      string    `json:"user"`
	Date_c      time.Time `json:"date"`
	Description string    `json:"description"`
	Phone       int64     `json:"phone"`
	Status      bool      `json:"status"`
}

// Clase que almacenara los datos obtenidos de cada tag.
type Tag struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
}

// Clase para almacenar tags de cada lugar
type Tagplace struct {
	Id_tags  string `json:"id_tag"`
	Id_place int64  `json:"id_place"`
}

//Horarios, 7 días de la semana, open-Close por cada día
type Schedule struct {
	Id int64     `json:"id"`
	O1 time.Time `json:"o1"`
	C1 time.Time `json:"c1"`
	O2 time.Time `json:"o2"`
	C2 time.Time `json:"c2"`
	O3 time.Time `json:"o3"`
	C3 time.Time `json:"c3"`
	O4 time.Time `json:"o4"`
	C4 time.Time `json:"c4"`
	O5 time.Time `json:"o5"`
	C5 time.Time `json:"c5"`
	O6 time.Time `json:"o6"`
	C6 time.Time `json:"c6"`
	O7 time.Time `json:"o7"`
	C7 time.Time `json:"c7"`
}

// Clase que permite almacenar evaluaciones.
type Evaluation struct {
	Id       int64     `json:"id"`
	Id_user  string    `json:"id_user"`
	Id_place string    `json:"id_place"`
	Score    int64     `json:"score"`
	Comment  string    `json:"comment"`
	Date     time.Time `json:"date"`
	Status   bool      `json:"status"`
}

// Clase que almacena coordenadas y id de un lugar.
type Geocoord struct {
	Id  int64   `json:"id"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
