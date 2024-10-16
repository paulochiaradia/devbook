package middlewares

import (
	"github.com/paulochiaradia/devbook/src/autenticacao"
	"github.com/paulochiaradia/devbook/src/respostas"
	"log"
	"net/http"
)

// Logger escreve informacoes sobre a rota chamada
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica se o usuario que fez a requisicao esta autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
