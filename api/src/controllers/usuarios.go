package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/paulochiaradia/devbook/src/autenticacao"
	"github.com/paulochiaradia/devbook/src/banco"
	"github.com/paulochiaradia/devbook/src/models"
	"github.com/paulochiaradia/devbook/src/repositorios"
	"github.com/paulochiaradia/devbook/src/respostas"
	"github.com/paulochiaradia/devbook/src/seguranca"
)

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario models.Usuario
	if erro := json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuario no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusFound, usuarios)
}

// BuscaUsuario busca um determinado usuario no banco de dados
func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	fmt.Println(parametro)

	usuarioId, erro := strconv.ParseUint(parametro["usuarioID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscaUsuarioID(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusFound, usuario)
}

// AtualizarUsuario atualiza dados de um usuario no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametro["usuarioID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("nao pode atualizar um usuario que nao seja o seu"))
		return
	}

	corpoDaRequisiscao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario models.Usuario
	if erro := json.Unmarshal(corpoDaRequisiscao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := usuario.Preparar("alterar"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.AtualizarUsuario(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletaUsuario deleta um usuario no banco de dados
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDnoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDnoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("nao pode excluir um usuario que nao seja o seu"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.DeletarUsuario(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
}

// SeguirUsuario o usuario da requisao segue o usuario parametro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioID == seguidorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("voce nao pode seguir voce mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.SeguirUsuario(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario o usuario da requisicao para de seguir o usario do parametro
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorId == usuarioId {
		respostas.Erro(w, http.StatusForbidden, errors.New("voce nao pode para de seguir voce mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.PararDeSeguirUsuario(usuarioId, seguidorId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores buscar todos os seguidores de um usuario
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer func(db *sql.DB) {
		erro := db.Close()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	seguindo, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguindo)
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioIDNoToken != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("nao e possivel alterar a senha de um ususario que nao seja o seu"))
		return
	}

	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var senha models.Senha
	if erro := json.Unmarshal(corpoDaRequisicao, &senha); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer func(db *sql.DB) {
		if erro := db.Close(); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}
