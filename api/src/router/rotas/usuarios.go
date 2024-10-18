package rotas

import (
	"net/http"

	"github.com/paulochiaradia/devbook/src/controllers"
)

var rotasUsuarios = []Rotas{

	//Adiciona um usuario no banco
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	//Busca todos os usuarios no banco
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	//Busca um determinado usuario no banco
	{
		URI:                "/usuario/{usuarioID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaUsuario,
		RequerAutenticacao: true,
	},
	//Atualiza um usuario no banco
	{
		URI:                "/usuario/{usuarioID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	//Exclui um usuario no banco
	{
		URI:                "/usuario/{usuarioID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletaUsuario,
		RequerAutenticacao: true,
	},
	//Usuario da requisao vai seguir o usuario do parametro
	{
		URI:                "/usuario/{usuarioID}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},

	//Usuario da requisicao vai parar de seguir usuario do parametro
	{
		URI:                "/usuario/{usuarioID}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},

	//Busca seguidores de um usuario
	{
		URI:                "/usuario/{usuarioID}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},

	//Busca quem o usuario segue
	{
		URI:                "/usuario/{usuarioID}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},

	// Atualiza a senha do usuario
	{
		URI:                "/usuario/{usuarioID}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
