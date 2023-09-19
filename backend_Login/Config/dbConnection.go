package config

import (
	"command/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var baseClientes *gorm.DB

func conexionDB() {
	var db_url = os.Getenv("DB_URL") //obtengo mi variable del .env
	if db_url == "" {
		panic("DB_URL environment variable missing")
	}
	var err error
	baseClientes, err = gorm.Open(postgres.Open(db_url), &gorm.Config{}) //abre la conexion

	if err != nil { //Si encontro un error
		panic(err)
	} else {
		fmt.Println("Conectado a la base de Datos")
	}
	autoMigracion(baseClientes)
}

// Migracioness a la base de datos
func autoMigracion(connection *gorm.DB) {
	//Modelos que quiero mirar
	connection.AutoMigrate(&models.Cliente{})
}
