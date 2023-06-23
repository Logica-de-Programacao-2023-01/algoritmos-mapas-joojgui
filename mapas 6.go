package main

import "fmt"

func somarContagens(contagens []map[string]int) map[string]int {
	soma := make(map[string]int)

	for _, contagem := range contagens {
		for palavra, quantidade := range contagem {
			soma[palavra] += quantidade
		}
	}

	return soma
}

func main() {

	contagens := []map[string]int{
		{"banana": 3, "maçã": 2},
		{"banana": 2, "laranja": 4},
		{"maçã": 1, "laranja": 1},
	}

	soma := somarContagens(contagens)

	fmt.Println("Soma das contagens:", soma)
}
