package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasMusicas = []Rota{ //Slice de struct rota para criar todas as rotas que envolvem usuário
	{
		Uri:                "/musicas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarMusica,
		RequerAutenticacao: false,
	}, // Rota para criar usuário
	/* {
		Uri:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	}, //rota para buscar todos os usuários
	{
		Uri:                "/usuarios/{usuarioid}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	}, //Rota para buscar um usuário em específico
	{
		Uri:                "/usuarios/{usuarioid}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarUsuario,
		RequerAutenticacao: true,
	}, //Rota para editar um usuário
	{
		Uri:                "/usuarios/{usuarioid}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	}, //Rota para deletar um usuário */
}
