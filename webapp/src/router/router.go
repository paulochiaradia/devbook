package router

import (
	"github.com/gorilla/mux"
	"github.com/paulochiaradia/devbook/webapp/src/router/rotas"
)

// Gerar gera o router que contem todas as rotas do projeto
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
