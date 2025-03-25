package addon

import "fmt"

// EXERCICIO 6

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func Produtos() {
	geleia := Produto{
		Nome:       "geleia",
		Preco:      10.99,
		Quantidade: 1,
	}

	chocolate := Produto{
		Nome:       "Chocolate",
		Preco:      8.99,
		Quantidade: 10,
	}

	Salgado := Produto{
		Nome:       "Salgado",
		Preco:      7.50,
		Quantidade: 3,
	}

	fmt.Println("\nNome:\t", Salgado.Nome, "\nPreço:\t", Salgado.Preco, "\nQuantidade: ", Salgado.Quantidade)
	fmt.Println("\nNome:\t", chocolate.Nome, "\nPreço:\t", chocolate.Preco, "\nQuantidade: ", chocolate.Quantidade)
	fmt.Println("\nNome:\t", geleia.Nome, "\nPreço:\t", geleia.Preco, "\nQuantidade: ", geleia.Quantidade)
}
