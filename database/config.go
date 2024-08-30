package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar variables de Entorno")
	}

	var err error
	username := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, pass, host, port, dbName)

	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatalf("Error al conectar con DB: %s", err)
	}

	fmt.Println("Conexion a DB Exitosa")
}
