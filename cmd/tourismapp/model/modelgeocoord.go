package model

import "log"

func QueryGeocoords() []Geocoord {
	connectDatabase()
	pingDatabase()
	geocoords := make([]Geocoord, 0)
	rows, errq := db.Query("SELECT * FROM geocoord")
	if errq != nil {
		log.Println(errq)
		return geocoords
	}
	defer rows.Close()

	//tmp almacena en cada iteración el objeto
	tmp := Geocoord{}

	for rows.Next() {
		err := rows.Scan(
			&tmp.Id,
			&tmp.Lat,
			&tmp.Lng)
		if err != nil {
			log.Println(err)
		}
		geocoords = append(geocoords, tmp)
	}
	disconnectDatabase()
	return geocoords
}

func QueryGeocoord(id string) []Geocoord {
	connectDatabase()
	pingDatabase()
	geocoord := make([]Geocoord, 0)
	geo := Geocoord{}
	stmt, errp := db.Prepare("SELECT * FROM geocoord WHERE id=$1")
	if errp != nil {
		log.Println("Error preparing query", errp)
		return geocoord
	}
	defer stmt.Close()
	errq := stmt.QueryRow(id).Scan(
		&geo.Id,
		&geo.Lat,
		&geo.Lng)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return geocoord
	}
	geocoord = append(geocoord, geo)
	return geocoord
}

func InsertGeocoord(geocoord Geocoord) bool {
	connectDatabase()
	pingDatabase()

	query, _ := db.Prepare("INSERT INTO geocoord VALUES ($1, $2, $3)")
	_, errq := query.Exec(
		geocoord.Id,
		geocoord.Lat,
		geocoord.Lng)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}

func UpdateGeocoord(geocoord Geocoord) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update geocoord set lat=$1, lng=$2 where id=$3")
	_, errq := query.Exec(
		geocoord.Lat,
		geocoord.Lng,
		geocoord.Id)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}
