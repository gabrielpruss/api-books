package database

import (
	"log"
	"time"

	//"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var user_db *gorm.DB

func StartDB() {
	//str := "host=localhost port:3306 user=admin dbname=livros.livros_fisicos sslmode=disable password="
	str := "root:123456@tcp(127.0.0.1:3306)/livros?charset=utf8mb4&parseTime=True&loc=Local"
	str_user := "root:123456@tcp(127.0.0.1:3306)/security?charset=utf8mb4&parseTime=True&loc=Local"

	/* db livros block */
	database, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	config, error := db.DB()
	if error != nil {
		log.Fatal("Error", error)
	}

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	/* end block*/

	/**************************************************************************************/

	/* db secutiry block */
	usr_database, usr_err := gorm.Open(mysql.Open(str_user), &gorm.Config{})
	if usr_err != nil {
		log.Fatal("Error: ", usr_err)
	}

	user_db = usr_database

	usr_config, usr_error := user_db.DB()
	if usr_error != nil {
		log.Fatal("Error", usr_error)
	}

	usr_config.SetMaxIdleConns(10)
	usr_config.SetMaxOpenConns(100)
	usr_config.SetConnMaxLifetime(time.Hour)
	/* end block*/

	/**************************************************************************************/

}

func GetDatabase() *gorm.DB {
	return db
}

func User_GetDatabase() *gorm.DB {
	return user_db
}
