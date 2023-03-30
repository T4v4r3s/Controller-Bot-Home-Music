package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	Uri:                "/login",
	Metodo:             http.MethodPost, // Geralmente usamos método POST para login!
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
