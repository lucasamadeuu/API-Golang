package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/lucasamadeuu/bookstore-API/handlers"
	"github.com/lucasamadeuu/bookstore-API/models"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", handlers.EncontrarLivros)
	router.POST("/books", handlers.CriarLivros)
	router.GET("/books/:id", handlers.EncontrarLivro)
	router.PATCH("/books/:id", handlers.AtualizarLivro)
	router.DELETE("/books/:id", handlers.DeletarLivro)

	router.Run()

}
