package controllers

import (
	"github.com/gabrielpruss/api-books/models"
	"github.com/gabrielpruss/api-books/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//db := database.User_GetDatabase()

	var p models.Login

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	// Autoriza usuario
	token, ret := services.AuthUser(p)
	switch ret {
	case -1:
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})
		services.RegisterLogin("", p.User, 0)
		return
	case -2:
		c.JSON(400, gin.H{
			"error": "invalid credentials",
		})
		services.RegisterLogin("", p.User, 0)
		return
	case -3:
		c.JSON(500, gin.H{
			"error": "cannot create token",
		})
		services.RegisterLogin("", p.User, 0)
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

	// Registra novo login
	services.RegisterLogin(token, p.User, 1)

}

/*
	var user models.User
	dbError := db.Where("user = ?", p.User).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})
		return
	}

	if user.Pwd != services.SHA256Encoder(p.Pwd) {
		c.JSON(400, gin.H{
			"error": "invalid credentials",
		})
		return
	}


	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
	"token": token,
	})
*/
