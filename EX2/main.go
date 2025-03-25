package main

import (
	"fmt"
	"addon/addon"
)

func First(){

	// Exercicio 1
	
		var ra [6]int

		ra[0] = 1
		ra[1] = 2
		ra[2] = 4
		ra[3] = 6
		ra[4] = 2
		ra[5] = 2

	//Exercicio 2
	
	fmt.Print(ra[0])
	fmt.Print(ra[1])
	fmt.Print(ra[2])
	fmt.Print(ra[3])
	fmt.Print(ra[4])
	fmt.Print(ra[5], "\n")
}

func Second(){
	//Exercicio 3
		
	var ex3[6] int

	for i := 0; i < len(ex3); i++{
		fmt.Printf("\nDigite o %dº número do RA:", i + 1)
		fmt.Scanln(&ex3[i])
	}
	for v := 0; v < len(ex3); v++{
		fmt.Printf("%d", ex3[v])
	}
	fmt.Printf("\n")
}

func Third(){
	//Exercicio 4
	
	var ex4 [6] int

	for index := range ex4 {
		fmt.Printf("Digite o %dº número do RA:", index + 1)
		fmt.Scanln(&ex4[index])
	}
	for value := range ex4{
		fmt.Printf("%d", ex4[value])
	}
	fmt.Printf("\n")
}


func main(){
	First() 		//Exercicio 1 e 2
	Second() 		//Exercicio 3
	Third()			//Exercicio 4
	addon.Matriz()	//Exercicio 5
	addon.Produtos()
}
