package routes

import (
	"gi-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	// HTML pages
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", controllers.ExibeIndexHtml)

	// API
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET(":nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.Run()
}
