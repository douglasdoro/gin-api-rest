package routes

import (
	"gi-api-rest/controllers"

	// docs "github.com/douglasdoro/gin-api-rest/docs"
	docs "github.com/douglasdoro/gin-api-rest/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func HandleRequest() {
	r := gin.Default()

	// Swagger
	docs.SwaggerInfo.BasePath = "/"

	// HTML pages
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", controllers.ExibeIndexHtml)
	r.GET("/list", controllers.ExibeListHtml)
	r.Static("/assets", "./assets")

	// Not Found
	r.NoRoute(controllers.RotaNaoEncontrada)

	// API
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET(":nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
