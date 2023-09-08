package services

import (
	"os/user"
	"time"

	"github.com/gabrielpruss/api-books/database"
	"github.com/gabrielpruss/api-books/models"
)

var iUser user.User

func AuthUser(login models.Login) (string, int) {
	db := database.User_GetDatabase()

	var user models.User
	dbError := db.Where("user = ?", login.User).First(&user).Error
	if dbError != nil {
		return "", -1
	}

	if user.Pwd != SHA256Encoder(login.Pwd) {
		return "", -2
	}

	token, ret := GetToken(user)
	if ret == 0 {
		return token, 0
	}

	token, err := NewJWTService().GenerateToken(user.ID)
	if err != nil {
		return "", -3
	}

	return token, 0
}

func GetToken(user models.User) (string, int) {
	db := database.User_GetDatabase()

	var reg models.Tokens
	dbError := db.Where("user = ?", user.User).Order("deadline DESC").First(&reg).Error
	if dbError != nil {
		return "", -1
	}

	timeNow := time.Now().Format("2006-01-02 15:04:05")
	if timeNow < reg.Deadline {
		return reg.Token, 0
	}

	return "", -1
}

/*
func RegisterToken(token, user string) {
	db := database.User_GetDatabase()

	var reg models.Tokens
	reg.Token = token
	reg.User = user
	reg.Dataproc = time.Now().Format("2006-01-02 15:04:05")
	reg.Deadline = time.Now().Add(2 * time.Hour).Format("2006-01-02 15:04:05")
	reg.Iporig = ""

	_ = db.Create(&reg).Error

	return
}*/

func RegisterLogin(token, user string, success uint) {
	db := database.User_GetDatabase()

	var reg models.Reclogin
	reg.Token = ""
	reg.User = user
	reg.Dataproc = time.Now().Format("2006-01-02 15:04:05")
	reg.Iporig = ""
	reg.Success = success

	_ = db.Create(&reg).Error

	if success == 1 {
		time := time.Now().Format("2006-01-02 15:04:05")
		_ = db.Model(&models.User{}).Where("user = ?", user).Updates(
			models.User{Lastlogin: time})
	}

	return
}
