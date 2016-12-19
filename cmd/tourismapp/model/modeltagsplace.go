package model

import (
	"log"
	"strconv"
)

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
	id_t, _ := idt
	id_p, _ := strconv.ParseInt(idp, 10, 8)
	errq := stmt.QueryRow(id_t, id_p).Scan(
		&t.Id_place,
		&t.Id_tags)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return tagplace
	}
	tagplace = append(tagplace, t)
	return tagplace
}

func InsertTagPlace(tagplace Tagplace) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("INSERT INTO tags_places VALUES ($1, $2)")
	_, errq := query.Exec(
		tagplace.Id_place,
		tagplace.Id_tags)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}
