package rotas

import (
	"github.com/paulochiaradia/devbook/src/controllers"
	"net/http"
)

// rotaLogin Login do usuario na API
var rotaLogin = Rotas{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
