package migrations

import (
	"github.com/gabrielpruss/api-books/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.Ebook)
	db.AutoMigrate(models.Livros_fisicos)
	db.AutoMigrate(models.SetLido)
	db.AutoMigrate(models.User)
	db.AutoMigrate(models.Login)
	db.AutoMigrate(models.Reclogin)
	db.AutoMigrate(models.Tokens)
}
