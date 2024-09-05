package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// produtos := []Produto{
	// 	{Nome: "Teste1", Descricao: "Descricao11", Preco: 39.12, Quantidade: 2},
	// 	{"Teste2", "Descricao22", 14, 5},
	// 	{"Teste33", "Nova Descricao", 11, 1},
	// }
	todosProdutos := sele

	temp.ExecuteTemplate(w, "index", produtos)
}
