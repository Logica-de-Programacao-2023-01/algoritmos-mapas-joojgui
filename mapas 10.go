package main

import "fmt"

func contarPares(slice []int) map[[2]int]int {
	frequencia := make(map[[2]int]int)

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			frequencia[pair]++
		}
	}

	return frequencia
}

func main() {
	// Exemplo de uso da função
	numeros := []int{1, 2, 3, 1, 2, 4, 5, 3, 4, 1, 5, 2, 3, 4}

	frequencia := contarPares(numeros)

	fmt.Println("Frequência de pares:")
	for par, freq := range frequencia {
		fmt.Printf("%v: %d\n", par, freq)
	}
}
