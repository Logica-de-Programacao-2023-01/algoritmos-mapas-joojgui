package main

import (
	"fmt"
	"strings"
)

func contarLetrasPorPalavra(frase string) map[string]map[rune]int {
	palavras := strings.Fields(frase) // Divide a frase em palavras

	contagem := make(map[string]map[rune]int)

	for _, palavra := range palavras {
		contagem[palavra] = contarLetras(palavra)
	}

	return contagem
}

func contarLetras(palavra string) map[rune]int {
	letras := make(map[rune]int)

	for _, letra := range palavra {
		letras[letra]++
	}

	return letras
}

func main() {

	frase := "O rato roeu a roupa do rei de Roma"

	contagem := contarLetrasPorPalavra(frase)

	for palavra, letras := range contagem {
		fmt.Println("Palavra:", palavra)
		for letra, quantidade := range letras {
			fmt.Printf("Letra: %c - Quantidade: %d\n", letra, quantidade)
		}
		fmt.Println()
	}
}
