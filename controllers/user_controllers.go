package controllers

import (
	"github.com/gabrielpruss/api-books/database"
	"github.com/gabrielpruss/api-books/models"
	"github.com/gabrielpruss/api-books/services"
	"github.com/gin-gonic/gin"
)

// Esta funcao cria um usuario na tabela security.users
func CreateUser(c *gin.Context) {
	db := database.User_GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	if len(user.Pwd) < 12 {
		c.JSON(400, gin.H{
			"error": "invalid password (min 12 caract): ",
		})
		return
	}

	user.Pwd = services.SHA256Encoder(user.Pwd)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
