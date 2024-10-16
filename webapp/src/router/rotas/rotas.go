package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota modelo de rota padrao
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar adiciona todas as rotas para um router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return router
}
