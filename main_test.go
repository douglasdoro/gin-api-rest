package main

import (
	"encoding/json"
	"gi-api-rest/controllers"
	"gi-api-rest/database"
	"gi-api-rest/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var alunoMockId int64

func setupDasRotasDeTeste() *gin.Engine {
	// Just for better visualization of tests
	gin.SetMode(gin.ReleaseMode)

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

func DeletaAlunoMock() {
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
	defer DeletaAlunoMock()

	r := setupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	defer DeletaAlunoMock()
	r := setupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := setupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	searchPath := "/alunos/" + strconv.Itoa(int(alunoMockId))
	req, _ := http.NewRequest("GET", searchPath, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do aluno teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678901", alunoMock.CPF, "Os CPF's devem ser iguais")
	assert.Equal(t, "123456789", alunoMock.RG, "Os RG's devem ser iguais")
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := setupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	deletePath := "/alunos/" + strconv.Itoa(int(alunoMockId))
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	respostaBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, `{"message":"Aluno deleted"}`, string(respostaBody))
}
