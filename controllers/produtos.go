package controllers

import (
	"net/http"
	"teste/models"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	// produtos := []Produto{
	// 	{Nome: "Teste1", Descricao: "Descricao11", Preco: 39.12, Quantidade: 2},
	// 	{"Teste2", "Descricao22", 14, 5},
	// 	{"Teste33", "Nova Descricao", 11, 1},
	// }
	todosProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}
