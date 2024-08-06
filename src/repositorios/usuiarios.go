package repositorios

import (
	"api/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar: Insere um usuário no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	var ultimoIdInserido uint64

	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values($1,$2,$3,$4) RETURNING id",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	erro = statement.QueryRow(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha).Scan(&ultimoIdInserido)
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// Atualizar: Atualiza informaçoes de ususário, com exceçao da senha que tera metodo proprio
func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = $1, nick = $2, email = $3 where id = $4",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil

}
