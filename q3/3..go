package main

import (
	"errors"
	"sort"
)

func calcularEstatisticas(lista []int) (int, int, float64, error) {
	if len(lista) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}

	// Ordenar a lista em ordem crescente
	sort.Ints(lista)

	// Encontrar o maior e o menor valor
	maiorValor := lista[len(lista)-1]
	menorValor := lista[0]

	// Calcular a média dos valores, desconsiderando o maior e o menor valor
	soma := 0
	for i := 1; i < len(lista)-1; i++ {
		soma += lista[i]
	}
	media := float64(soma) / float64(len(lista)-2)

	return maiorValor, menorValor, media, nil
}
