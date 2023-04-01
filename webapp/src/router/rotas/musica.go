package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasMusica = []Rota{
	{
		URI:                "/musica",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeAdicionarMusicas,
		RequerAutenticacao: false,
	},
	{
		URI:                "/addmusica",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AddMusica,
		RequerAutenticacao: false,
	},
}
