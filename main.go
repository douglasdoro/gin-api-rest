package main

import (
	"gi-api-rest/database"
	"gi-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
