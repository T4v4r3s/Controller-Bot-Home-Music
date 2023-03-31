package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// Insere os templates HTML na variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html")) //carrega qualquer coisa .html
}

// Executa e renderiza uma página html
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
