package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates carrega todos os templates para a variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("./views/*.html"))
}

// ExecutarTemplate executa/renderiza um template na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	for _, tmpl := range templates.Templates() {
		fmt.Println(tmpl.Name())
	}
	templates.ExecuteTemplate(w, template, dados)

}
