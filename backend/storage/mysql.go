package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Importamos el driver MySQL
)

// Db es la variable global para la conexión a la base de datos
var Db *sql.DB

// InitDB inicializa la conexión con la base de datos
func InitDB() {
	var err error
	// Cambia esta cadena con tu URL de conexión a la base de datos
	connStr := "root:password@tcp(localhost:3306)/recordatorios?parseTime=true"
	Db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}

	// Verificamos que la conexión se haya establecido correctamente
	err = Db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos: ", err)
	}

	fmt.Println("Conectado a la base de datos")
}
