package model

import (
	"database/sql"

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
		log.Fatalln("Error opening database: %q", err)
		DATABASE_URL := "postgres://eozcyemimcuhgg:3ac2YMMZ0EMofFw6rdrTXIky6W@ec2-107-22-250-212.compute-1.amazonaws.com:5432/da6rnltctu258a"
		db, err = sql.Open("postgres", DATABASE_URL)
	}
}

func disconnectDatabase() {
	err = db.Close()
	if err != nil {
		log.Fatalln("Error closing database: %q", err)
	}
}

func pingDatabase() {
	err = db.Ping()
	if err != nil {
		log.Fatalln("Error ping to database", err)
	}
}

// Responses method GET, all data
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
			log.Fatal(err)
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
			&tmp.Descripcion)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
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
			&tmp.pos)
		if err != nil {
			log.Fatal(err)
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
			&tmp.o1,
			&tmp.c1,
			&tmp.o2,
			&tmp.c2,
			&tmp.o3,
			&tmp.c3,
			&tmp.o4,
			&tmp.c4,
			&tmp.o5,
			&tmp.c5,
			&tmp.o6,
			&tmp.c6,
			&tmp.o7,
			&tmp.c7)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
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
			log.Fatal(err)
		}
		tagsplaces = append(tagsplaces, tmp)
	}
	disconnectDatabase()
	return tagsplaces
}

// Responses methods GET, one data
func QueryCustomer(rut string) []Customer {
	connectDatabase()
	pingDatabase()
	customer := make([]Customer, 0)
	cus := Customer{}
	errq := db.QueryRow("SELECT * FROM customer WHERE rut=$1", rut).Scan(
		&cus.Name,
		&cus.Surname,
		&cus.S_surname,
		&cus.Rut,
		&cus.Mail,
		&cus.Password)
	disconnectDatabase()
	if errq != nil {
		log.Fatalln("Error in query ", errq)
		return customer
	}
	customer = append(customer, cus)
	return customer
}

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
