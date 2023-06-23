package main

import "fmt"

func gerarSequenciaFibonacci(n int) map[int]int {
	sequencia := make(map[int]int)

	a, b := 0, 1

	for i := 0; a <= n; i++ {
		sequencia[i] = a
		a, b = b, a+b
	}

	return sequencia
}

func main() {

	n := 100

	sequencia := gerarSequenciaFibonacci(n)

	fmt.Println("Sequência de Fibonacci até", n, ":")
	for i, num := range sequencia {
		fmt.Printf("%d: %d\n", i, num)
	}
}
