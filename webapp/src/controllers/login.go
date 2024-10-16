package controllers

import (
	"net/http"

	"github.com/paulochiaradia/devbook/webapp/src/utils"
)

// CarregarTelaDeLogin renderiza a tela de login da aplicacao
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
