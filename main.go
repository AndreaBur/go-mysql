package main

import (
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// creamos una instancia de contacto
	newContact := models.Contact{
		Name:  "Nuevo Usuario",
		Email: "araña@newemail.com",
		Phone: "3107774532",
	}
	handlers.CreateContact(db, newContact)

	handlers.ListContacts(db)

	//handlers.GetContactByID(db, 3)

}
