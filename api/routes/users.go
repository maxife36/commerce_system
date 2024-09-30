package routes

import (
	// "commerce-system/database/dao/category"
	"commerce-system/api/handlers"
	"commerce-system/database/repository"
	"commerce-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.RouterGroup) {

	user := r.Group("/users")

	user.GET("/", handlers.GetAllUsers)
	user.POST("/", testEndpoint)
}

func allUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Ingresaste a Ususarios",
	})
}

func testEndpoint(c *gin.Context) {

	//CREAR UN REGISTRO

	// var category models.Category

	// if err := c.BindJSON(&category); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// result, err := dao.CreateCategory(&category)

	// LEER TODOS LOS REGISTROS

	// result, err := dao.GetAllCategories()

	// LEER UN REGISTRO POR SU ID

	// idStr := c.Param("id")

	// id, err := strToUint(idStr)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// result, err := category.GetById(id)

	//BUSCAR POR PARAMETRO (QUERY PARAMS)

	category := repository.NewDao(models.Category{})

	if details := c.Query("details"); details == "true" {
		category.Details = true
	}

	result, err := category.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
