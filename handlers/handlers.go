package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {
	//mi funcion recibe la conexion (db *sql.DB); procedemos a hacer la consulta porque ya estamos conectados
	query := "SELECT * FROM contact"
	//despues de la consulta, la ejecutamos
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//vamos a iterar sobre los resultados y mostrarlos
	fmt.Println("\n LISTA DE CONTACTOS :")
	fmt.Println("\n -------------------------------------------------------------------")
	for rows.Next() {
		//instancia de modelo contact
		contact := models.Contact{}
		/*
			//si hay valor nulo, en este caso se hace el ejemplo con "email", se tiene:

			var valueEmail sql.NullString
			err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
			if err != nil {
				log.Fatal(err)
			}
			//verifico los valores
			if valueEmail.Valid{
				contact.Email = valueEmail.String
			}else{
				contact.Email = "Sin correo electrónico"
			}

		*/

		err := rows.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, NOMBRE: %s, E-MAIL: %s, TEL: %s \n", contact.Id, contact.Name,
			contact.Email, contact.Phone)
		fmt.Println("\n ---------------------------------------------------------------")

	}
}

// creamos la funcion "GetContactByID" que obtiene un contacto de la base de datos mediante el ID
func GetContactByID(db *sql.DB, contactID int) {
	//consulta SQL para seleccionar un contacto por su ID
	query := "SELECT * FROM contact WHERE Id = ?"
	row := db.QueryRow(query, contactID)
	//hacemos una instancia del modelo contact
	contact := models.Contact{}
	/*
		//si tenemos un valor nulo, creamos la variable para darle manejo
		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
				if err != nil {
					if err == sql.ErrNoRows {
						log.Fatal("No se encontró nungún contacto con el ID : %d", contactID)
					}
				}
				//verifico los valores
				if valueEmail.Valid{
					contact.Email = valueEmail.String
				}else{
					contact.Email = "Sin correo electrónico"
				}

	*/

	err := row.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No se encontró nungún contacto con el ID : %d", contactID)
		}
	}

	fmt.Println("\n LISTA DE CONTACTOS :")
	fmt.Println("\n -------------------------------------------------------------------")
	fmt.Printf("ID: %d, NOMBRE: %s, E-MAIL: %s, TEL: %s \n", contact.Id, contact.Name,
		contact.Email, contact.Phone)
	fmt.Println("\n ---------------------------------------------------------------")

}
