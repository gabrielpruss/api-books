package controllers

import (
	"fmt"
	"net/http"

	"github.com/gabrielpruss/api-books/database"
	"github.com/gabrielpruss/api-books/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
 * Aqui estao as funcoes de api chamadas via router
 * Cada funcao tem seu proposito e sua URL API especifica
 * Atentar-se para os parametros da URL
 */

// Esta funcao retorna uma relacao com todos ebooks da tabela livros.ebooks
func ShowEbook(c *gin.Context) {
	db := database.GetDatabase()

	var ebooks []models.Ebook
	err := db.Find(&ebooks).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})
		return
	}
	c.JSON(200, ebooks)
}

// Esta funcao retorna uma relacao de todos ebooks de um determinado autor
func ShowEbooksAutor(c *gin.Context) {
	tmp := c.Param("autor")
	autor := "%" + tmp + "%"

	db := database.GetDatabase()

	var ebooks []models.Ebook
	//err := db.Find(&fisicos).Error
	err := db.Where("autor LIKE ?", autor).Find(&ebooks).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})
		return
	}
	c.JSON(200, ebooks)
}

// Esta funcao retorna as informacoes de um ou mais livros pelo titulo
func ShowEbooksTitulo(c *gin.Context) {
	tmp := c.Param("titulo")
	titulo := "%" + tmp + "%"

	db := database.GetDatabase()

	var ebooks []models.Ebook
	//err := db.Find(&fisicos).Error
	err := db.Where("titulo LIKE ?", titulo).Find(&ebooks).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})
		return
	}
	c.JSON(200, ebooks)
}

// Esta funcao retorna uma relacao de todos ebooks lidos
func ShowEbooksLidos(c *gin.Context) {
	db := database.GetDatabase()

	var ebooks []models.Ebook
	//err := db.Find(&fisicos).Error
	err := db.Where("lido = 1").Find(&ebooks).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})
		return
	}
	c.JSON(200, ebooks)
}

// Esta funcao cria um ebook na tabela livros.ebooks
func CreateEbook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Ebook

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create ebook: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

// Essa funcao seta um ebook como lido - se houver data tambem
func SetLidoEbook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.SetLido

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON" + err.Error(),
		})
		return
	}

	titulo := "%" + book.Titulo + "%"
	autor := "%" + book.Autor + "%"
	dataproc := book.Dataproc
	pages := book.Paginas

	result := db.Model(&models.Ebook{}).Where("titulo LIKE ? AND autor LIKE ?", titulo, autor).Updates(
		models.Ebook{Lido: 1, Dataproc: dataproc, Paginas: pages})
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot set status" + err.Error(),
		})
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "Book not found",
		})
		return
	}

	ack := BuscaEbookPeloTitulo(c, db, titulo)
	if ack == 0 {
		c.JSON(200, book)
	}
}

// busca um ebook pelo titulo
func BuscaEbookPeloTitulo(c *gin.Context, db *gorm.DB, titulo string) int {

	var ebook models.Ebook
	err := db.Where("titulo LIKE ?", titulo).First(&ebook).Error
	if err != nil {
		return 0
	}
	c.JSON(200, ebook)
	return 1
}

// Esta funcao retorna uma relacao com todos livros fisicos da tabela livros.livros_fisicos
func ShowFisicos(c *gin.Context) {
	db := database.GetDatabase()

	var fisicos []models.Livros_fisicos
	err := db.Find(&fisicos).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})
		return
	}
	c.JSON(200, fisicos)
}

// Esta funcao retorna uma relacao de todos livros fisicos de um determinado autor
func ShowFisicosAutor(c *gin.Context) {
	tmp := c.Param("autor")
	autor := "%" + tmp + "%"

	db := database.GetDatabase()

	var fisicos []models.Livros_fisicos
	//err := db.Find(&fisicos).Error
	err := db.Where("autor LIKE ?", autor).Find(&fisicos).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})
		return
	}
	c.JSON(200, fisicos)
}

// Esta funcao retorna uma relacao de todos livros fisicos com determinado titulo
func ShowFisicosTitulo(c *gin.Context) {
	tmp := c.Param("titulo")
	titulo := "%" + tmp + "%"

	db := database.GetDatabase()

	var fisicos []models.Livros_fisicos
	//err := db.Find(&fisicos).Error
	err := db.Where("titulo LIKE ?", titulo).Find(&fisicos).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})
		return
	}
	c.JSON(200, fisicos)
}

// Esta funcao retorna uma relacao de todos livros fisicos lidos
func ShowFisicosLidos(c *gin.Context) {
	db := database.GetDatabase()

	var fisicos []models.Livros_fisicos
	//err := db.Find(&fisicos).Error
	err := db.Where("lido = 1").Find(&fisicos).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})
		return
	}
	c.JSON(200, fisicos)
}

// Esta funcao cria um livro na tabela livros.livros_fisicos
func CreateFisico(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Livros_fisicos

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	//TODO: validar a existencia de todos os campos necessarios: titulo/autor

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create ebook: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

// Esta funcao set um livro fisico como lido - se houver data tambem
func SetLidoFisico(c *gin.Context) {
	db := database.GetDatabase()

	var book models.SetLido

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON" + err.Error(),
		})
		return
	}

	titulo := "%" + book.Titulo + "%"
	autor := "%" + book.Autor + "%"
	dataproc := book.Dataproc
	pages := book.Paginas

	result := db.Model(&models.Livros_fisicos{}).Where("titulo LIKE ? AND autor LIKE ?", titulo, autor).Updates(
		models.Livros_fisicos{Lido: 1, Dataproc: dataproc, Paginas: pages})
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot set status" + err.Error(),
		})
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "Book not found",
		})
		return
	}

	ack := BuscaFisicoPeloTitulo(c, db, titulo)
	if ack == 0 {
		c.JSON(200, book)
	}
}

func BuscaFisicoPeloTitulo(c *gin.Context, db *gorm.DB, titulo string) int {
	var fisico models.Livros_fisicos

	err := db.Where("titulo LIKE ?", titulo).First(&fisico).Error
	if err != nil {
		return 0
	}
	c.JSON(200, fisico)
	return 1
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
