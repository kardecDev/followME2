package banco

import (
	"api/config"
	"database/sql"

	_ "github.com/lib/pq" //Driver
)

// Conectar: Execua a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("postgres", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
