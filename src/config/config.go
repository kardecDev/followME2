package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco: É onde vamos configurar a string de conexao com banco de dados
	StringConexaoBanco = ""
	//Porta_DB: É onde amos configurar a porta de servico do banco de dados
	Porta_DB = 0

	///Porta_API: É onde amos configurar a porta de servico da api
	Porta_API = 0
)

// Carregar: Vai inicializar as variaves de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta_API, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		Porta_API = 5000 //Deault da api
	}

	Porta_DB, erro = strconv.Atoi(os.Getenv("DB_PORT"))
	if erro != nil {
		Porta_DB = 5432 //Deault do Postgres
	}

	StringConexaoBanco = os.Getenv("DATABASE_URL")
}
