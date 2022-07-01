package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasamadeuu/bookstore-API/models"
)

type CriarLivrosInput struct {
	Titulo string `json:"titulo" binding:"required"`
	Autor  string `json:"autor" binding:"required"`
}

type AtualizarLivroInput struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func EncontrarLivros(c *gin.Context) {
	var livros []models.Livro
	models.DB.Find(&livros)

	c.JSON(http.StatusOK, gin.H{"data": livros})
}

func CriarLivros(c *gin.Context) {
	var input CriarLivrosInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err.Error()})
		return
	}

	livro := models.Livro{
		Titulo: input.Titulo,
		Autor:  input.Autor,
	}
	models.DB.Create(&livro)

	c.JSON(http.StatusOK, gin.H{"data": livro})
}

func EncontrarLivro(c *gin.Context) {

	var livro models.Livro
	err := models.DB.Where("id = ?", c.Param("id")).First(&livro).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Livro não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": livro})
}
