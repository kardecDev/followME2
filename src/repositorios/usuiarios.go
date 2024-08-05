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

// Busca usuário por ID
func (repositorio usuarios) BuscarUsuarioPorID(id uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = $1",
		id,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}
