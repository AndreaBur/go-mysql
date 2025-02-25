package handlers

import (
	"database/sql"
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
}
//vamos a iterar sobre los resultados y mostrarlos
fmt.Println("\n LISTA DE CONTACTOS :")
fmt.Println("\n -------------------------------------------------------------------")
for rows.Next(){
	//instancia de modelo contact
	contact := models.Contact{}
	err := rows.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("ID: %d, NOMBRE: %s, E-MAIL: %s, TEL: %s \n", contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("\n ---------------------------------------------------------------")

}