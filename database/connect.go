package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	//ponemos el dns que nos permite conctar este archivo con mi base de datos en mysql
	//con godotenv.load carga todas las variables de entorno
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	//dns pero con variables de entorno
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))

	//abrimos una conexion con la base de datos
	//Open devuelve la conexion de la base pero tambien un error
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	//verificamos si se mantiene la conexion
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión a la base de datos MySql exitosa")
	return db, nil
}
