package controllers

import (
	"gi-api-rest/models"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos1(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Douglas",
	})
}

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo bem?",
	})
}
