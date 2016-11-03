package model

import "time"

//import "database/sql/driver"

// Clase que almacenara los datos obtenidos de cada usuario.
type Customer struct {
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	S_surname string `json:"s_surname" binding:"required"`
	Rut       string `json:"rut"`
	Mail      string `json:"mail" binding:"required"`
	Password  string `json:"pass" binding:"required"`
}

// Clase que almacenara los datos obtenidos de cada lugar.
type Place struct {
	Id          int64
	Name        string
	Score       int64
	User_c      string
	Date_c      time.Time
	Descripcion string
}

// Clase que almacenara los datos obtenidos de cada tag.
type Tag struct {
	Id   int64
	Name string
}

// Clase para almacenar tags de cada lugar
type Tagplace struct {
	Id_tags  int64
	Id_place int64
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
	Id       int64
	Id_user  string
	Id_place string
	Score    int64
	Comment  string
	Date     time.Time
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
