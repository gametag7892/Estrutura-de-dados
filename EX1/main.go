package main

import "fmt"

func ParOuImpar(n int)(string){
	if n%2 == 0{
		return "É par"
	}else{
		return "É impar"
	}
}

func SomaAte(n int)(int){
	var soma int
	for i := 0; i <= n; i++{
		soma = soma + i
	}
	return soma
}

func ClassificarNota(nota int)(string){
	switch {
		case nota >= 9 && nota <= 10:
			return "Excelente"
		case nota >= 7 && nota <= 8:
			return "Bom"
		case nota >= 5 && nota <= 6:
			return "Regular"
		case nota >= 3 && nota <= 4:
			return "Ruim"
		case nota <= 2 && nota >= 0:
			return "Muito ruim"
		default:
			return "digite um número de 0 a 10"
	}
}

func main(){
	fmt.Println(ParOuImpar(11))
	fmt.Println(SomaAte(4))
	fmt.Println(ClassificarNota(2))
}