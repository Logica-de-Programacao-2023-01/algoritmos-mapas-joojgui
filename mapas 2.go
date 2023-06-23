package main

import "fmt"

func mesclarMapas(mapa1, mapa2 map[string]string) map[string]string {
	resultado := make(map[string]string)

	// Copia os elementos do mapa1 para o resultado
	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}

	// Adiciona ou atualiza os elementos do mapa2 no resultado
	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado
}

func main() {
	// Exemplo de uso da função
	mapa1 := map[string]string{
		"chave1": "valor1",
		"chave2": "valor2",
	}

	mapa2 := map[string]string{
		"chave2": "novoValor",
		"chave3": "valor3",
	}

	resultado := mesclarMapas(mapa1, mapa2)

	fmt.Println("Resultado:", resultado)
}
