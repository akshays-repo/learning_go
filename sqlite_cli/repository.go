package main

import (
	"database/sql"
	"log"
)

func addNewPerson(db *sql.DB, newPerson person) bool {
	stmt, _ := db.Prepare("INSERT INTO people (id, first_name, last_name, email, ip_address) VALUES (?, ?, ?, ?, ?)")
	stmt.Exec(nil, newPerson.first_name, newPerson.last_name, newPerson.email, newPerson.ip_address)
	defer stmt.Close()
	log.Fatal("Added", newPerson)
	return true

}
