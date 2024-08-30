package routes

import "github.com/gin-gonic/gin"

func UsersRoutes(r *gin.RouterGroup) {

	user := r.Group("/users")

	user.GET("/", allUsers)
}

func allUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Ingresaste a Ususarios",
	})
}
