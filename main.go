package main

import (
	"gi-api-rest/models"
	"gi-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Douglas", CPF: "00000000", RG: "0000000"},
		{Nome: "Ana", CPF: "00000000", RG: "0000000"},
	}

	routes.HandleRequest()
}
