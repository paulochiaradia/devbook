package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/paulochiaradia/devbook/src/autenticacao"
	"github.com/paulochiaradia/devbook/src/banco"
	"github.com/paulochiaradia/devbook/src/models"
	"github.com/paulochiaradia/devbook/src/repositorios"
	"github.com/paulochiaradia/devbook/src/respostas"
	"github.com/paulochiaradia/devbook/src/seguranca"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioNoBanco, erro := repositorio.BucarUsuarioEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if erro := seguranca.VerificarSenha(usuarioNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, _ := autenticacao.CriarToken(usuarioNoBanco.ID)
	usuarioID := strconv.FormatUint(usuarioNoBanco.ID, 10)
	respostas.JSON(w, http.StatusOK, models.DadosAutenticacao{ID: usuarioID, Token: token})
}
