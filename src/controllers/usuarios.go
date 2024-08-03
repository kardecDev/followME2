package controllers

import (
	"api/banco"
	"api/modelos"
	"api/repositorios"
	"api/respostas"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CriarUsuario: Chama repositorio para criar um ususario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuario: Retorna uma lista contendo todos os usuarios da aplicacao
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	buscaUsuario := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(buscaUsuario)

	respostas.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario: Retorna um usuario especifico atraves de chave
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuario"))
}

// AtualizaUsuario: Atualiza informacoes de um usuario especifico atraves de chave
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alterar usuario"))
}

// DeletaUsuario: Remove um usario especifico atraves de chave
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluir usuario"))
}
