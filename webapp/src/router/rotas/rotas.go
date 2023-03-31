package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Representa as rotas da aplucação
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas no router
func Configurar(router *mux.Router) *mux.Router {

	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotasMusica...)

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	//Configura o fileserver para que ele fique no /assets/
	fileserver := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileserver))

	return router

}
