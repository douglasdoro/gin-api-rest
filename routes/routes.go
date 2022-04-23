package routes

import (
	"gi-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET(":nome", controllers.Saudacao)
	r.Run()
}
