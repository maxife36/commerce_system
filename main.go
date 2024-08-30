package main

import (
	"commerce-system/api/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Genero Router
	r := gin.Default()

	//Cargo varibales de entorno
	if err := godotenv.Load(); err != nil {
		fmt.Println("Not .env loaded")
	}

	//agrupo rutas
	app := r.Group("/api")

	routes.UsersRoutes(app)

	//Recupero Variable de Entorno y levanto servidor
	ADDR := os.Getenv("SERVER_ADDRESS")
	log.Fatalln(r.Run(ADDR))
}
