package main

import "fmt"

func somarValores(mapa map[string]int) int {
	soma := 0

	for _, valor := range mapa {
		soma += valor
	}

	return soma
}

func main() {
	// Exemplo de uso da função
	valores := map[string]int{
		"valor1": 10,
		"valor2": 20,
		"valor3": 30,
	}

	total := somarValores(valores)

	fmt.Println("Soma dos valores:", total)
}
