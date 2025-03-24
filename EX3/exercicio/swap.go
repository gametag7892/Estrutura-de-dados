package exercicio

import "fmt"

func Pont(){
	var b float64 = 0.2
	var p2* float64 = &b
	fmt.Println("Valor antes da modificação: ", b)
	*p2 = 2.2
	fmt.Println("Valor após a modificação: ", *p2)
}

func Swap(a, b int){
	var p1, p2 *int = &a, &b
	var temp *int

	temp = p1
	p1 = p2
	p2 = temp

	fmt.Print("antes do swap:\n")
	fmt.Println("x = ", a)
	fmt.Println("y = ", b)

	fmt.Print("depois do swap:\n")
	fmt.Println("x = ", *p1)
	fmt.Println("y = ", *p2)
}

