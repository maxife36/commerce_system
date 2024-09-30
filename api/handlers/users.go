package handlers

import (
	"commerce-system/database/repository"
	"commerce-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	userDao := repository.NewDao(models.User{})

	if details := c.Query("details"); details == "true" {
		userDao.Details = true
	}

	result, err := userDao.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
