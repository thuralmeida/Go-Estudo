package models

import (
	"teste/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConexaoBancoDados()

	selectTodosProdutos, err := db.Query("Select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			continue
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func BuscaProduto(idProduto string) Produto {
	db := db.ConexaoBancoDados()

	selectProduto, err := db.Query("Select * from produtos where id = $1 limit 1", idProduto)
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}

	for selectProduto.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			continue
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}

	defer db.Close()
	return p
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConexaoBancoDados()

	insereDadosBanco, err := db.Prepare("Insert Into Produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConexaoBancoDados()

	deletaProdutoBanco, err := db.Prepare("Delete From Produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletaProdutoBanco.Exec(id)
	defer db.Close()
}

func AtualizaProduto(nome, descricao string, preco float64, quantidade, id int) {
	db := db.ConexaoBancoDados()

	atualizaProdutoBanco, err := db.Prepare("update Produtos set nome=$1, descricao=$2, quantidade=$3, preco=$4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	atualizaProdutoBanco.Exec(nome, descricao, quantidade, preco, id)
	defer db.Close()
}
