package model

import "log"

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

func UpdateTag(tag Tag) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update tag set name=$1 where id=$2")
	_, errq := query.Exec(
		tag.Name,
		tag.Id)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}
