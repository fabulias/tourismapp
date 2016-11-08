package model

import "log"

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

func InsertSchedule(schedule Schedule) bool {
	connectDatabase()
	pingDatabase()

	query, _ := db.Prepare("INSERT INTO schedule VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)")
	_, errq := query.Exec(
		schedule.Id,
		schedule.O1,
		schedule.C1,
		schedule.O2,
		schedule.C2,
		schedule.O3,
		schedule.C3,
		schedule.O4,
		schedule.C4,
		schedule.O5,
		schedule.C5,
		schedule.O6,
		schedule.C6,
		schedule.O7,
		schedule.C7)

	disconnectDatabase()
	//log.Fatalln(errq)
	if errq != nil {
		return false
	} else {
		return true
	}
}

func UpdateSchedule(schedule Schedule) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update schedule set o1=$1,c1=$2,o2 =$3,c2 =$4,o3 =$5,c3 =$6,o4 =$7,c4 =$8,o5 =$9,c5=$10,o6 =$11,c6 =$12,o7 =$13,c7 =$14 where id=$15")
	_, errq := query.Exec(
		schedule.O1,
		schedule.C1,
		schedule.O2,
		schedule.C2,
		schedule.O3,
		schedule.C3,
		schedule.O4,
		schedule.C4,
		schedule.O5,
		schedule.C5,
		schedule.O6,
		schedule.C6,
		schedule.O7,
		schedule.C7,
		schedule.Id)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}
