package main

import (
	"gi-api-rest/controllers"
	"gi-api-rest/database"
	"gi-api-rest/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var alunoMockId int64

func setupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Nome do aluno teste",
		CPF:  "12345678901",
		RG:   "123456789",
	}

	database.DB.Create(&aluno)
	alunoMockId = int64(aluno.ID)
}

func DeletaAluno() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, alunoMockId)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := setupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/douglas", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	// }

	assert.Equal(t, http.StatusOK, resposta.Code, "They should be equal")

	mockDaResposta := `{"API diz":"E ai douglas, tudo bem?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAluno()

	r := setupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}
