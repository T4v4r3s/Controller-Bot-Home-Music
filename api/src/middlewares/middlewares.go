package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

// Escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc { //Next = próxima função
	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		next(w, r)
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r) //executar depois de validado
	}

}
