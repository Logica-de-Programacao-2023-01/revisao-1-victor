package main

import (
	"errors"
	"strings"
)

func calcularMediaLetrasPorPalavra(texto string) (float64, error) {
	palavras := strings.Fields(texto)
	totalPalavras := len(palavras)
	totalLetras := 0

	if totalPalavras == 0 {
		return 0, errors.New("texto vazio")
	}

	for _, palavra := range palavras {
		totalLetras += len(palavra)
	}

	media := float64(totalLetras) / float64(totalPalavras)
	return media, nil
}
