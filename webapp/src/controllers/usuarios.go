package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/paulochiaradia/devbook/webapp/src/config"
	"github.com/paulochiaradia/devbook/webapp/src/respostas"
)

// CriarUsuario chama a API para cadastrar um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	resposta, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
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

	//fmt.Println(resposta.StatusCode, resposta.Request)

	if resposta.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, resposta)
		return
	}

	respostas.JSON(w, resposta.StatusCode, nil)

}
