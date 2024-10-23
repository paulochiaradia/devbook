package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/paulochiaradia/devbook/webapp/src/config"
	"github.com/paulochiaradia/devbook/webapp/src/cookies"
	"github.com/paulochiaradia/devbook/webapp/src/modelos"
	"github.com/paulochiaradia/devbook/webapp/src/requisicoes"
	"github.com/paulochiaradia/devbook/webapp/src/respostas"

	"github.com/paulochiaradia/devbook/webapp/src/utils"
)

// CarregarTelaDeLogin renderiza a tela de login da aplicacao
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastro renderiza a pagina de cadastro de usuario
func CarregarPaginaDeCadastro(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal renderiza a pagina de cadastro de usuario
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	resposta, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
			return
		}
	}(resposta.Body)

	if resposta.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, resposta)
		return
	}

	var publicacoes []modelos.Publicacao
	if erro := json.NewDecoder(resposta.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

func CarregarPaginaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoID"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoId)
	resposta, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
			return
		}
	}(resposta.Body)

	if resposta.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, resposta)
		return
	}

	var publicacao modelos.Publicacao
	if erro := json.NewDecoder(resposta.Body).Decode(&publicacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)

}

func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, nomeOuNick)
	resposta, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
			return
		}
	}(resposta.Body)

	if resposta.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, resposta)
		return
	}

	var usuarios []modelos.Usuario

	if erro := json.NewDecoder(resposta.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

func CarregarPerfilDeUsuarios(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if usuarioID == usuarioLogadoID {
		http.Redirect(w, r, "/perfil", http.StatusFound)
		return
	}

	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         modelos.Usuario
		UsuarioLogadoID uint64
	}{
		Usuario:         usuario,
		UsuarioLogadoID: usuarioLogadoID,
	})

}

func CarregarPerfilDeUsuarioLogado(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", usuario)
}

func CarregarPaginaDePerfilDeUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	canal := make(chan modelos.Usuario)
	go modelos.BuscarDadosDoUsuario(canal, usuarioID, r)
	usuario := <-canal

	if usuario.ID == 0 {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao buscar usuario"})
		return
	}

	utils.ExecutarTemplate(w, "editar-usuario.html", usuario)
}
