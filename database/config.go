package database

import (
	"commerce-system/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func migrate() {

	err := db.AutoMigrate(
		&models.Brand{},
		&models.Category{},
		&models.Client{},
		&models.Combo{},
		&models.ComboDetails{},
		&models.Product{},
		&models.Production{},
		&models.ProductionDetails{},
		&models.Purchase{},
		&models.PurchaseDetails{},
		&models.Role{},
		&models.Sale{},
		&models.SaleDetails{},
		&models.Supplier{},
		&models.User{},
	)

	if err != nil {
		log.Fatalln("Se produjo un error en las migraciones de modelos a la DB")
	}

	log.Println("Migracion exitosa")
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar variables de Entorno")
	}

	var err error
	username := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, pass, host, port, dbName)

	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatalf("Error al conectar con DB: %s", err)
	}

	fmt.Println("Conexion a DB Exitosa")

	migrate()
}
