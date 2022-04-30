package main

import (
	"github.com/douglasdoro/gin-api-rest/database"
	"github.com/douglasdoro/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
