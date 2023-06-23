package main

import "fmt"

func calcularDivisaoGastos(gastos map[string]float64) map[string]float64 {
	total := 0.0
	numPessoas := len(gastos)

	for _, valor := range gastos {
		total += valor
	}

	media := total / float64(numPessoas)

	divisao := make(map[string]float64)
	for nome, valor := range gastos {
		divisao[nome] = valor - media
	}

	return divisao
}

func main() {

	gastos := map[string]float64{
		"Alice": 50.0,
		"Bob":   70.0,
		"Carol": 30.0,
		"David": 60.0,
		"Eve":   20.0,
	}

	divisao := calcularDivisaoGastos(gastos)

	fmt.Println("Divisão de gastos:")
	for nome, valor := range divisao {
		fmt.Printf("%s: %.2f\n", nome, valor)
	}
}

