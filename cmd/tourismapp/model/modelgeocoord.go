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

func QueryGeocoordRadio(lat, lng, radio string) []Geocoord {
	connectDatabase()
	pingDatabase()
	geocoords := make([]Geocoord, 0)

	query := "SELECT id, lat, lng" +
		" FROM (" +
		" SELECT id, lat, lng, ( 6371 * acos(cos(radians($1)) * cos(radians(lat)) * cos(radians(lng) - radians($2)) + sin(radians($1)) * sin(radians(lat)))) AS distance" +
		" FROM geocoord" +
		" ORDER BY distance" +
		" ) dynamic_t" +
		" WHERE distance < $3;"

	stmt_q, errp_q := db.Prepare(query)

	if errp_q != nil {
		log.Println("Error preparing query", errp_q)
		return geocoords
	}

	defer stmt_q.Close()
	rows, errq_e := stmt_q.Query(lat, lng, radio) //geo.Lat, geo.Lng, geo.Lat, radio)

	if errq_e != nil {
		log.Println(errq_e)
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
