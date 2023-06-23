package main

import (
	"fmt"
	"strings"
)

func contarPalavras(texto string) map[string]int {
	ocorrencias := make(map[string]int)
	palavras := strings.Fields(texto)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {

	texto := "Olá mundo olá novamente mundo"

	resultado := contarPalavras(texto)

	for palavra, ocorrencias := range resultado {
		fmt.Printf("Palavra: %s | Ocorrências: %d\n", palavra, ocorrencias)
	}
}
