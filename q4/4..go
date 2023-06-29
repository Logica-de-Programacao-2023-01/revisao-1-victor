package main

import (
	"errors"
)

func calcularPrecoFinal(precoBase float64, estadoDestino string, codigoImposto int) (float64, error) {
	if precoBase <= 0 {
		return 0, errors.New("preço base inválido")
	}

	var aliquotaImposto float64
	switch codigoImposto {
	case 1:
		aliquotaImposto = 0.05
	case 2:
		aliquotaImposto = 0.1
	case 3:
		aliquotaImposto = 0.15
	default:
		return 0, errors.New("imposto não encontrado")
	}

	var percentualFrete float64
	switch estadoDestino {
	case "SP":
		percentualFrete = 0.1
	case "RJ":
		percentualFrete = 0.15
	case "MG":
		percentualFrete = 0.2
	case "ES":
		percentualFrete = 0.25
	default:
		percentualFrete = 0.3
	}

	precoFinal := precoBase + (precoBase * aliquotaImposto) + (precoBase * percentualFrete)
	return precoFinal, nil
}
