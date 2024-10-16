package rotas

import (
	"github.com/paulochiaradia/devbook/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rotas{
	// Criar uma publicacao
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},

	// Buscar todas as publicacoes
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	// Busca uma publicacao
	{
		URI:                "/publicacoes/{publicacaoID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	// Atualiza uma publicacao
	{
		URI:                "/publicacoes/{publicacaoID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizaPublicacao,
		RequerAutenticacao: true,
	},

	// Deleta uma publicacao
	{
		URI:                "/publicacoes/{publicacaoID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},

	// Busca todas as publicacoes de um determinado usuario
	{
		URI:                "/usuarios/{usuarioID}/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticacao: true,
	},

	// Curtir uma publicacao
	{
		URI:                "/publicacoes/{publicacaoID}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},

	// Curtir uma publicacao
	{
		URI:                "/publicacoes/{publicacaoID}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
