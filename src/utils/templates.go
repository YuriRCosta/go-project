package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates insere os templates html na variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutarTemplate renderiza um template html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
