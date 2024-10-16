package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/paulochiaradia/devbook/src/autenticacao"
	"github.com/paulochiaradia/devbook/src/banco"
	"github.com/paulochiaradia/devbook/src/models"
	"github.com/paulochiaradia/devbook/src/repositorios"
	"github.com/paulochiaradia/devbook/src/respostas"
	"io"
	"net/http"
	"strconv"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioIdNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var publicacao models.Publicacao
	if erro := json.Unmarshal(corpoDaRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioIdNoToken

	if erro := publicacao.Preparar(); erro != nil {
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

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.CriarPublicacao(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametro["publicacaoID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, nil)
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

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPublicacao(publicacaoId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusFound, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
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

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.BuscarPublicacoes(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

func AtualizaPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoID"], 10, 64)
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
		if erro := db.Close(); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPublicacao(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("nao e possivel atualizar uma publicacao que nao e sua"))
		return
	}

	corpoDaRequiicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro := json.Unmarshal(corpoDaRequiicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := repositorio.AtualizarPublicacao(publicacaoID, publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoID"], 10, 64)
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
		if erro := db.Close(); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPublicacao(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("nao e possivel deletar uma publicacao que nao e sua"))
		return
	}

	if erro := repositorio.DeletarPublicacao(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametro["usuarioID"], 10, 64)
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
		if erro := db.Close(); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.BuscarPublicacoesPorUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoID"], 10, 64)
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
		if erro := db.Close(); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	if erro := repositorio.CurtirPublicacao(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoID"], 10, 64)
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
		if erro := db.Close(); erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}(db)

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	if erro := repositorio.DescurtirPublicacao(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}
