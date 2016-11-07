package model

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
	"log"
)

var (
	err error
	db  *sql.DB = nil
)

func connectDatabase() {
	db, err = sql.Open("postgres", "postgres://eozcyemimcuhgg:3ac2YMMZ0EMofFw6rdrTXIky6W@ec2-107-22-250-212.compute-1.amazonaws.com:5432/da6rnltctu258a") //os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Error opening database: %q", err)
		DATABASE_URL := "postgres://eozcyemimcuhgg:3ac2YMMZ0EMofFw6rdrTXIky6W@ec2-107-22-250-212.compute-1.amazonaws.com:5432/da6rnltctu258a"
		db, err = sql.Open("postgres", DATABASE_URL)
	}
}

func disconnectDatabase() {
	err = db.Close()
	if err != nil {
		log.Println("Error closing database: %q", err)
	}
}

func pingDatabase() {
	err = db.Ping()
	if err != nil {
		log.Println("Error ping to database", err)
	}
}

//
// Responses method GET, all data
//
func QueryCustomers() []Customer {
	connectDatabase()
	pingDatabase()
	customers := make([]Customer, 0)
	rows, errq := db.Query("SELECT * FROM customer")
	if errq != nil {
		log.Println(errq)
		return customers
	}
	defer rows.Close()

	//tmp almacena en cada iteración el objeto
	tmp := Customer{}
	for rows.Next() {
		err := rows.Scan(
			&tmp.Name,
			&tmp.Surname,
			&tmp.S_surname,
			&tmp.Rut,
			&tmp.Mail,
			&tmp.Password)
		if err != nil {
			log.Println(err)
		}
		customers = append(customers, tmp)
	}
	disconnectDatabase()
	return customers
}

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
			&tmp.Descripcion,
			&tmp.Fono)
		if err != nil {
			log.Println(err)
		}
		places = append(places, tmp)
	}
	disconnectDatabase()
	return places
}

func QueryEvaluations() []Evaluation {
	connectDatabase()
	pingDatabase()
	evaluations := make([]Evaluation, 0)
	rows, errq := db.Query("SELECT * FROM evaluation")
	if errq != nil {
		log.Println(errq)
		return evaluations
	}
	defer rows.Close()

	//tmp almacena en cada iteración el objeto
	tmp := Evaluation{}
	for rows.Next() {
		err := rows.Scan(
			&tmp.Id,
			&tmp.Id_user,
			&tmp.Id_place,
			&tmp.Score,
			&tmp.Comment,
			&tmp.Date)
		if err != nil {
			log.Println(err)
		}
		evaluations = append(evaluations, tmp)
	}
	disconnectDatabase()
	return evaluations
}

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
			&tmp.Id_place,
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

func QuerySchedules() []Schedule {
	connectDatabase()
	pingDatabase()
	schedules := make([]Schedule, 0)
	rows, errq := db.Query("SELECT * FROM schedule")
	if errq != nil {
		log.Println(errq)
		return schedules
	}
	defer rows.Close()

	//tmp almacena en cada iteración el objeto
	tmp := Schedule{}
	for rows.Next() {
		err := rows.Scan(
			&tmp.Id,
			&tmp.O1,
			&tmp.C1,
			&tmp.O2,
			&tmp.C2,
			&tmp.O3,
			&tmp.C3,
			&tmp.O4,
			&tmp.C4,
			&tmp.O5,
			&tmp.C5,
			&tmp.O6,
			&tmp.C6,
			&tmp.O7,
			&tmp.C7)
		if err != nil {
			log.Println(err)
		}
		schedules = append(schedules, tmp)
	}
	disconnectDatabase()
	return schedules
}

func QueryTags() []Tag {
	connectDatabase()
	pingDatabase()
	tags := make([]Tag, 0)
	rows, errq := db.Query("SELECT * FROM tags")
	if errq != nil {
		log.Println(errq)
		return tags
	}
	defer rows.Close()

	//tmp almacena en cada iteración el objeto
	tmp := Tag{}
	for rows.Next() {
		err := rows.Scan(
			&tmp.Id,
			&tmp.Name)
		if err != nil {
			log.Println(err)
		}
		tags = append(tags, tmp)
	}
	disconnectDatabase()
	return tags
}

func QueryTagsPlaces() []Tagplace {
	connectDatabase()
	pingDatabase()
	tagsplaces := make([]Tagplace, 0)
	rows, errq := db.Query("SELECT * FROM tags_places")
	if errq != nil {
		log.Println(errq)
		return tagsplaces
	}
	defer rows.Close()

	//tmp almacena en cada iteración el objeto
	tmp := Tagplace{}
	for rows.Next() {
		err := rows.Scan(
			&tmp.Id_tags,
			&tmp.Id_place)
		if err != nil {
			log.Println(err)
		}
		tagsplaces = append(tagsplaces, tmp)
	}
	disconnectDatabase()
	return tagsplaces
}

//
// Responses methods GET, one data
//
func QueryCustomer(rut string) []Customer {
	connectDatabase()
	pingDatabase()
	customer := make([]Customer, 0)
	cus := Customer{}
	stmt, errp := db.Prepare("SELECT * FROM customer WHERE rut=$1")
	if errp != nil {
		log.Println("Error preparing query", errp)
		return customer
	}
	defer stmt.Close()
	errq := stmt.QueryRow(rut).Scan(
		&cus.Name,
		&cus.Surname,
		&cus.S_surname,
		&cus.Rut,
		&cus.Mail,
		&cus.Password)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return customer
	}
	customer = append(customer, cus)
	return customer
}

func QueryEvaluation(id string) []Evaluation {
	connectDatabase()
	pingDatabase()
	evaluation := make([]Evaluation, 0)
	ev := Evaluation{}
	stmt, errp := db.Prepare("SELECT * FROM evaluation WHERE id=$1")
	if errp != nil {
		log.Println("Error preparing query", errp)
		return evaluation
	}
	defer stmt.Close()
	errq := stmt.QueryRow(id).Scan(
		&ev.Id,
		&ev.Id_user,
		&ev.Id_place,
		&ev.Score,
		&ev.Comment,
		&ev.Date)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return evaluation
	}
	evaluation = append(evaluation, ev)
	return evaluation
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
		&geo.Id_place,
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

func QueryPlace(id string) []Place {
	connectDatabase()
	pingDatabase()
	place := make([]Place, 0)
	plc := Place{}
	stmt, errp := db.Prepare("SELECT * FROM place WHERE id=$1")
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
		&plc.Descripcion,
		&plc.Fono)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return place
	}
	place = append(place, plc)
	return place
}

func QuerySchedule(rut string) []Schedule {
	connectDatabase()
	pingDatabase()
	schedule := make([]Schedule, 0)
	sh := Schedule{}
	stmt, errp := db.Prepare("SELECT * FROM schedule WHERE id=$1")
	if errp != nil {
		log.Println("Error preparing query", errp)
		return schedule
	}
	defer stmt.Close()
	errq := stmt.QueryRow(rut).Scan(
		&sh.Id,
		&sh.O1,
		&sh.C1,
		&sh.O2,
		&sh.C2,
		&sh.O3,
		&sh.C3,
		&sh.O4,
		&sh.C4,
		&sh.O5,
		&sh.C5,
		&sh.O6,
		&sh.C6,
		&sh.O7,
		&sh.C7)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return schedule
	}
	schedule = append(schedule, sh)
	return schedule
}

func QueryTag(id string) []Tag {
	connectDatabase()
	pingDatabase()
	tag := make([]Tag, 0)
	t := Tag{}
	stmt, errp := db.Prepare("SELECT * FROM tags WHERE id=$1")
	if errp != nil {
		log.Println("Error preparing query", errp)
		return tag
	}
	defer stmt.Close()
	errq := stmt.QueryRow(id).Scan(
		&t.Id,
		&t.Name)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return tag
	}
	tag = append(tag, t)
	return tag
}

func QueryTagsPlace(idp, idt string) []Tagplace {
	connectDatabase()
	pingDatabase()
	tagplace := make([]Tagplace, 0)
	t := Tagplace{}
	stmt, errp := db.Prepare("SELECT * FROM tags_places WHERE id_tag=$1 AND id_place=$2")
	if errp != nil {
		log.Println("Error preparing query", errp)
		return tagplace
	}
	defer stmt.Close()
	id_t, _ := strconv.ParseInt(idt, 10, 8)
	id_p, _ := strconv.ParseInt(idp, 10, 8)
	errq := stmt.QueryRow(id_t, id_p).Scan(
		&t.Id_tags,
		&t.Id_place)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return tagplace
	}
	tagplace = append(tagplace, t)
	return tagplace
}

//Post methods
//Inserta un nuevo customer en la base de datos, retorna true si el insert fue
// exitoso, false si ya existe en la bdd
func InsertCustomer(user Customer) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("INSERT INTO customer VALUES ($1,$2,$3,$4,$5,$6)")
	_, errq := query.Exec(
		user.Name,
		user.Surname,
		user.S_surname,
		user.Rut,
		user.Mail,
		user.Password)
	disconnectDatabase()

	if errq != nil {
		return false
	} else {
		return true
	}
}

//Inserta un nuevo Place en la base de datos, retorna true o false dependiendo
//del exito de la operación
func InsertPlace(place Place) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("INSERT INTO place VALUES ($1,$2,$3,$4,$5,$6)")
	_, errq := query.Exec(
		place.Id,
		place.Name,
		place.Score,
		place.User_c,
		place.Date_c,
		place.Descripcion)
	disconnectDatabase()

	if errq != nil {
		return false
	} else {
		return true
	}
}
func InsertTag(tag Tag) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("INSERT INTO tags VALUES ($1, $2)")
	_, errq := query.Exec(
		tag.Id,
		tag.Name)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}

func InsertTagPlace(tagplace Tagplace) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("INSERT INTO tags_places VALUES ($1, $2)")
	_, errq := query.Exec(
		tagplace.Id_tags,
		tagplace.Id_place)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}

func InsertSchedule(schedule Schedule) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("INSERT INTO schedule VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $10, $11, $12, $13, $14, $15)")
	_, errq := query.Exec(schedule.Id, schedule.O1, schedule.C1, schedule.O2,
		schedule.C2, schedule.O3, schedule.C3, schedule.O4,
		schedule.C4, schedule.O6, schedule.C6, schedule.O7,
		schedule.C7)

	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}

func InsertEvaluation(evaluation Evaluation) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("INSERT INTO evaluation VALUES ($1, $2, $3, $4, $5, $6)")
	_, errq := query.Exec(
		evaluation.Id,
		evaluation.Id_user,
		evaluation.Id_place,
		evaluation.Score,
		evaluation.Comment,
		evaluation.Date)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}
