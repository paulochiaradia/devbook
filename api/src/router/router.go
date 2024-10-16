package router

import (
	"github.com/gorilla/mux"
	"github.com/paulochiaradia/devbook/src/router/rotas"
)

func GerarRouter() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
