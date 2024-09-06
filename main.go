package main

import (
	"commerce-system/api/routes"
	"commerce-system/database"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Conexion a la DB
	database.Connect()
	//AutoMigarte Models
	database.Migarte()
	//Genero Router
	r := gin.Default()

	//agrupo rutas
	app := r.Group("/api")

	routes.UsersRoutes(app)

	//Cargo varibales de entorno
	if err := godotenv.Load(); err != nil {
		fmt.Println("Not .env loaded")
	}
	//Recupero Variable de Entorno y levanto servidor
	ADDR := os.Getenv("SERVER_ADDRESS")
	log.Fatalln(r.Run(ADDR))
}
