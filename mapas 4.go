package main

import (
	"fmt"
	"sort"
	"strings"
)

func agruparAnagramas(palavras []string) map[string][]string {
	grupos := make(map[string][]string)

	for _, palavra := range palavras {
		chave := ordenarCaracteres(palavra)
		grupos[chave] = append(grupos[chave], palavra)
	}

	return grupos
}

func ordenarCaracteres(palavra string) string {

	palavra = strings.ToLower(palavra)

	caracteres := strings.Split(palavra, "")

	sort.Strings(caracteres)

	return strings.Join(caracteres, "")
}

func main() {

	palavras := []string{"amor", "roma", "mar", "cão", "não", "roma", "ar", "moça"}

	grupos := agruparAnagramas(palavras)

	for _, grupo := range grupos {
		fmt.Printf("Anagramas de %s: %v\n", grupo[0], grupo)
	}
}
