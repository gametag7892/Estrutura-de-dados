package main
import (
	"fmt"
)

// Definição da pilha
type Stack struct {
	items []rune
}

// Push adiciona um caractere na pilha
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop remove e retorna o caractere do topo da pilha
func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return ' ', false
	}

	// Pega o último elemento
	ultimoIndice := len(s.items) - 1
	item := s.items[ultimoIndice]

	// Remove da pilha
	s.items = s.items[:ultimoIndice]
	return item, true
}

// Função para verificar se uma palavra é um palíndromo usando pilha
func isPalindrome(word string) bool {
	var stack Stack

	// Empilha todos os caracteres
	for _, letra := range word {
		stack.Push(letra)
	}

	// Constrói a string invertida
	var reverso []rune
	for range word {
		if letra, ok := stack.Pop(); ok {
			reverso = append(reverso, letra)
		}
	}

	// Compara com a original
	return word == string(reverso)
}
func main() {
	// Testes
	words := []string{"radar", "arara", "golang", "level", "hello"}

	for _, word := range words {
		if isPalindrome(word) {
			fmt.Printf("\"%s\" é um palíndromo\n", word)
		} else {
			fmt.Printf("\"%s\" NÃO é um palíndromo\n", word)
		}
	}
}