package repositorios

import (
	"api/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar: Insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

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

<<<<<<< HEAD
// Buscar: Tras todos os usuarios que satisfação ao filtro de nome e nick
func (repositorio Usuarios) Buscar(buscaUsuario string) ([]modelos.Usuario, error) {

	buscaUsuario = fmt.Sprintf("%%%s%%", buscaUsuario)

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoem from usuarios where nome like $1 or nick like $2",
		buscaUsuario, buscaUsuario)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil

}

// Busca usuário por ID
func (repositorio Usuarios) BuscarUsuarioPorID(id uint64) (modelos.Usuario, error) {
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

// Atualizar: Atualiza informaçoes de ususário, com exceçao da senha que tera metodo proprio
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
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
=======
>>>>>>> main
