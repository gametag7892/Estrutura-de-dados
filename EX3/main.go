package main

import (
	"exercicio/exercicio"
	"fmt"
)


func main(){
	var a, b = 1, 2

	exercicio.Pont()
	fmt.Println("")
	exercicio.Swap(a, b)

	lista := exercicio.TaskList{}

	lista.AddTask("Estudar go")
	lista.AddTask("Fazer compras")
	lista.AddTask("Academia")

	fmt.Println("\nLista encadeada:")
	fmt.Println("")
	lista.ShowTask()
	fmt.Println("")

	lista.CompleteTask()
	lista.ShowTask()
	fmt.Println("")

	lista.CompleteTask()
	lista.ShowTask()
	fmt.Println("")
}