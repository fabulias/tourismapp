package model

import "time"

//import "database/sql/driver"

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
	o1 time.Time
	c1 time.Time
	o2 time.Time
	c2 time.Time
	o3 time.Time
	c3 time.Time
	o4 time.Time
	c4 time.Time
	o5 time.Time
	c5 time.Time
	o6 time.Time
	c6 time.Time
	o7 time.Time
	c7 time.Time
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
	pos      Point
}

// Clase que almacena las coordenadas.
type Point struct {
	Lat float64
	Lng float64
}

//func (u *Point) Scan(value interface{}) error {
//	*u = Point(value.(float64))
//	return nil
//}
//func (u Point) Value() (driver.Value, error) {
//	return Point(u), nil
//}
