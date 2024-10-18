package rotas

import (
	"github.com/paulochiaradia/devbook/webapp/src/controllers"
	"net/http"
)

var rotaPaginaPrincipal = Rota{
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
