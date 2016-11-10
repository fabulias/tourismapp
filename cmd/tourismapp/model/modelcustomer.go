package model

import "log"

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
			&tmp.Password,
			&tmp.Status)
		if err != nil {
			log.Println(err)
		}
		customers = append(customers, tmp)
	}
	disconnectDatabase()
	return customers
}

func QueryCustomer(rut string) []Customer {
	connectDatabase()
	pingDatabase()
	customer := make([]Customer, 0)
	cus := Customer{}
	stmt, errp := db.Prepare("SELECT * FROM customer WHERE rut=$1 AND status=true")
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
		&cus.Password,
		&cus.Status)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return customer
	}
	customer = append(customer, cus)
	return customer
}

func QueryCustomerM(mail string) []Customer {
	connectDatabase()
	pingDatabase()
	customer := make([]Customer, 0)
	cus := Customer{}
	stmt, errp := db.Prepare("SELECT * FROM customer WHERE mail=$1 AND status=true")
	if errp != nil {
		log.Println("Error preparing query", errp)
		return customer
	}
	defer stmt.Close()
	errq := stmt.QueryRow(mail).Scan(
		&cus.Name,
		&cus.Surname,
		&cus.S_surname,
		&cus.Rut,
		&cus.Mail,
		&cus.Password,
		&cus.Status)
	disconnectDatabase()
	if errq != nil {
		log.Println("Error in query ", errq)
		return customer
	}
	customer = append(customer, cus)
	return customer
}

//Inserta un nuevo customer en la base de datos, retorna true si el insert fue
// exitoso, false si ya existe en la bdd
func InsertCustomer(user Customer) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("INSERT INTO customer VALUES ($1,$2,$3,$4,$5,$6,$7)")
	_, errq := query.Exec(
		user.Name,
		user.Surname,
		user.S_surname,
		user.Rut,
		user.Mail,
		user.Password,
		user.Status)
	disconnectDatabase()

	if errq != nil {
		return false
	} else {
		return true
	}
}

func UpdateCustomer(customer Customer) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update customer set name=$1, surname=$2, s_surname=$3, mail=$4, pass=$5, status=$6 where rut=$7")
	_, errq := query.Exec(
		customer.Name,
		customer.Surname,
		customer.S_surname,
		customer.Mail,
		customer.Password,
		customer.Status,
		customer.Rut)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}

func EraseCustomer(rut string) bool {
	connectDatabase()
	pingDatabase()
	query, _ := db.Prepare("update customer set status='false' where rut=$1")
	_, errq := query.Exec(rut)
	disconnectDatabase()
	if errq != nil {
		return false
	} else {
		return true
	}
}
