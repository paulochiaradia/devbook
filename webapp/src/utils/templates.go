package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates carrega todos os templates para a variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("./views/*.html"))
	templates = template.Must(templates.ParseGlob("./views/templates/*.html"))
}

// ExecutarTemplate executa/renderiza um template na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
