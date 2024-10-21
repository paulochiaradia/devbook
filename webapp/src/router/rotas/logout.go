package rotas

import (
	"net/http"

	"github.com/paulochiaradia/devbook/webapp/src/controllers"
)

var rotaLogout = Rota{
	URI:                "/logout",
	Metodo:             http.MethodGet,
	Funcao:             controllers.FazerLogout,
	RequerAutenticacao: false,
}
