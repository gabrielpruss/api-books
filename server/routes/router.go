package routes

import (
	"log"
	"net/http"

	"github.com/gabrielpruss/api-books/controllers"
	"github.com/gabrielpruss/api-books/server/middlewares"
	"github.com/gin-gonic/gin"
)

// nome das rotas URL e seus respectivos parametros
func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		/***************************** EBOOKS ******************************/
		books := main.Group("ebooks", middlewares.Auth())
		{
			books.GET("/", controllers.ShowEbook)
			books.POST("/", controllers.CreateEbook)
		}
		books = main.Group("ebooks_lidos", middlewares.Auth())
		{
			books.GET("/", controllers.ShowEbooksLidos)
			books.POST("/", controllers.SetLidoEbook)
		}
		books = main.Group("ebooks_autor", middlewares.Auth())
		{
			books.GET("/:autor", controllers.ShowEbooksAutor)
		}
		books = main.Group("ebooks_titulo", middlewares.Auth())
		{
			books.GET("/:titulo", controllers.ShowEbooksTitulo)
		}
		/*******************************************************************/

		/***************************** FISICOS *****************************/
		books = main.Group("livros_fisicos", middlewares.Auth())
		{
			books.GET("/", controllers.ShowFisicos)
			books.POST("/", controllers.CreateFisico)
		}
		books = main.Group("fisicos_lidos", middlewares.Auth())
		{
			books.GET("/", controllers.ShowFisicosLidos)
			books.POST("/", controllers.SetLidoFisico)
		}
		books = main.Group("fisicos_autor", middlewares.Auth())
		{
			books.GET("/:autor", controllers.ShowFisicosAutor)
		}
		books = main.Group("fisicos_titulo", middlewares.Auth())
		{
			books.GET("/:titulo", controllers.ShowFisicosTitulo)
		}
		/*******************************************************************/

		/***************************** LOGINS ******************************/
		route := main.Group("user", middlewares.Auth())
		{
			route.POST("/", controllers.CreateUser)
		}
		route = main.Group("login")
		{
			route.POST("/", controllers.Login)
		}
		/*******************************************************************/

	}

	return router
}

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
