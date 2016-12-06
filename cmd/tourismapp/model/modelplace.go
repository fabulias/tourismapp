package model

import "log"

func QueryPlaces() []Place {
	connectDatabase()
	pingDatabase()
	places := make([]Place, 0)
	rows, errq := db.Query("SELECT * FROM place")
	if errq != nil {
		log.Println(errq)
		return places
	}
	defer rows.Close()

	//tmp almacena en cada iteración el objeto
	tmp := Place{}
	for rows.Next() {
		err := rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.Score,
			&tmp.User_c,
			&tmp.Date_c,
			&tmp.Description,
			&tmp.Phone,
			&tmp.Status)
		if err != nil {
			log.Println(err)
		}
		places = append(places, tmp)
	}
	disconnectDatabase()
	return places
}

func QueryPlace(id string) []Place {
	connectDatabase()
	pingDatabase()
	place := make([]Place, 0)
	plc := Place{}
	stmt, errp := db.Prepare("SELECT * FROM place WHERE id=$1 AND status=true")
	if errp != nil {
		log.Println("Error preparing query", errp)
		return place
	}
	defer stmt.Close()
	errq := stmt.QueryRow(id).Scan(
		&plc.Id,
		&plc.Name,
		&plc.Score,
		&plc.User_c,
		&plc.Date_c,
		&plc.Description,
		&plc.Phone,
		&plc.Status)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return place
	}
	place = append(place, plc)
	return place
}

//Inserta un nuevo Place en la base de datos, retorna true o false dependiendo
//del exito de la operación
func get_id() int64 {
	connectDatabase()
	pingDatabase()
	var id_last int64
	row, errq := db.Query("SELECT count(*) FROM place")
	if errq != nil {
		log.Println(errq)
	}
	defer row.Close()
	for row.Next() {

		err := row.Scan(&id_last)
		if err != nil {
			log.Println(err)
			log.Println(row)
			log.Println("HOLA")
		}
	}
	id_last++
	disconnectDatabase()
	return id_last

}
func InsertPlace(place Place) bool {
	connectDatabase()
	pingDatabase()
	score_default := 0
	query, _ := db.Prepare("INSERT INTO place (name, score, id_user, date_inscription, description, phone, status)VALUES ($1,$2,$3,$4,$5,$6,$7)")
	_, errq := query.Exec(
		place.Name,
		score_default,
		place.User_c,
		place.Date_c,
		place.Description,
		place.Phone,
		true)
	disconnectDatabase()

	if errq != nil {
		log.Println(errq)
		return false
	} else {
		return true
	}
}

func UpdatePlace(place Place) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update place set name=$1, score=$2, user_c=$3, date_c=$4, description=$5, phone=$6 where id=$7")
	_, errq := query.Exec(
		place.Name,
		place.Score,
		place.User_c,
		place.Date_c,
		place.Description,
		place.Status,
		place.Id)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}

func ErasePlace(id string) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update place set status='false' where id=$1")
	_, errq := query.Exec(id)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}
