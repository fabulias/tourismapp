package model

import "log"

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
			&tmp.Date,
			&tmp.Status)
		if err != nil {
			log.Println(err)
		}
		evaluations = append(evaluations, tmp)
	}
	disconnectDatabase()
	return evaluations
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
		&ev.Date,
		&ev.Status)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return evaluation
	}
	evaluation = append(evaluation, ev)
	return evaluation
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
		evaluation.Date,
		evaluation.Status)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}

func UpdateEvaluation(eval Evaluation) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update evaluation set id_user=$1, id_place=$2, score=$3, comment=$4, date=$5, status=$6 where id=$7")
	_, errq := query.Exec(
		eval.Id_user,
		eval.Id_place,
		eval.Score,
		eval.Comment,
		eval.Date,
		eval.Status,
		eval.Id)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}

func EraseEvaluation(id string) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update customer set status='false' where rut=$1")
	_, errq := query.Exec(id)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}
