package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//Aqui ficam todas as funções que carregam páginas!

// CarregarTelaDeLogin carrega tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil) //Carrega a página de login não passando dados para ela!
}

// CarregarTelaDeLogin carrega tela de criação de usuarios
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

func CarregarPaginaDeAdicionarMusicas(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "addmusica.html", nil)
}
