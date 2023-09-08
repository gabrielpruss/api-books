package main

import (
	"fmt"

	"github.com/gabrielpruss/api-books/database"
	"github.com/gabrielpruss/api-books/server"
	"github.com/gabrielpruss/api-books/website"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	database.StartDB()
	fmt.Println("Iniciando o servidor Rest com Go")
	//routes.HandleRequest()

	server := server.NewServer()

	website.Run()

	server.Run()
}
