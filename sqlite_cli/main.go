package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dixonwille/wmenu"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./names.db")
	defer db.Close()
	checkErr(err)
	menu := wmenu.NewMenu("What would you like to do?")
	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })
	menu.Option("Add a new Person", 0, true, nil)
	menu.Option("Find a Person", 1, false, nil)
	menu.Option("Update a Person's information", 2, false, nil)
	menu.Option("Delete a person by ID", 3, false, nil)
	menuerr := menu.Run()
	if menuerr != nil {
		log.Fatal(menuerr)
	}

}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {
	switch opts[0].Value {
	case 0:
		newPerson := handleNewPerson()
		addNewPerson(db, newPerson)
		break

	case 1:
		fmt.Println("Finding a Person")
	case 2:
		fmt.Println("Update a Person's information")
	case 3:
		fmt.Println("Deleting a person by ID")
	case 4:
		fmt.Println("Quitting application")
	}
}

func handleNewPerson() person {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a first name ")
	firstName, _ := reader.ReadString('\n')

	fmt.Print("Enter a last name ")
	lastName, _ := reader.ReadString('\n')

	fmt.Print("Enter a email ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Enter a IP address ")
	ipAddress, _ := reader.ReadString('\n')

	newPerson := person{
		first_name: firstName,
		last_name:  lastName,
		email:      email,
		ip_address: ipAddress,
	}

	return newPerson
}
