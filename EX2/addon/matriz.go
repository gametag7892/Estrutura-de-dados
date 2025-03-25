package addon

import (
	"fmt"
)

// EXERCICIO 5

func Matriz() { 
	var alunos[3][3] string;

	for i := range alunos {
		fmt.Printf("Digite o índice do aluno %d:", i + 1)
		fmt.Scanln(&alunos[i][0])
		fmt.Printf("Digite o RA do aluno %d:", i + 1)
		fmt.Scanln(&alunos[i][1])
		fmt.Printf("Digite o Nome do aluno %d:", i + 1)
		fmt.Scanln(&alunos[i][2])
	} 

	fmt.Printf("\nDados dos Alunos:\n")
	fmt.Printf("--------------------")
	fmt.Print("\nIndice	RA	Nome\n")

	for i := range alunos {
		for j := range alunos[i] {
			fmt.Print(alunos[i][j],"\t")
		}
		fmt.Println()
	} 

	
	
}
