package main

import (
	"commerce-system/api/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var addr string

func init() {

	//Cargo varibales de entorno
	if err := godotenv.Load(); err != nil {
		fmt.Println("Not .env loaded")
	}
	//Recupero Variable de Entorno
	addr = os.Getenv("SERVER_ADDRESS")
}

func main() {
	//Genero Router
	r := gin.Default()

	//agrupo rutas
	app := r.Group("/api")

	routes.UsersRoutes(app)

	log.Fatalln(r.Run(addr))
}
