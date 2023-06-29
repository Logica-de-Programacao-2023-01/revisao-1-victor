package main

import (
	"errors"
)

type Cliente struct {
	Compras []float64
}

func verificarElegibilidadeDesconto(compraAtual float64, cliente Cliente) (float64, error) {
	if compraAtual <= 0 {
		return 0, errors.New("valor da compra inválido")
	}

	totalCompras := calcularTotalCompras(cliente)
	mediaCompras := calcularMediaCompras(totalCompras, cliente)

	desconto := 0.0

	if totalCompras > 1000 {
		desconto = compraAtual * 0.1
	} else if totalCompras <= 1000 {
		desconto = compraAtual * 0.05
	} else if totalCompras <= 500 {
		desconto = compraAtual * 0.02
	}

	if len(cliente.Compras) == 0 {
		desconto = compraAtual * 0.1
	}

	if mediaCompras > 1000 {
		desconto = compraAtual * 0.2
	}

	return desconto, nil
}

func calcularTotalCompras(cliente Cliente) float64 {
	total := 0.0
	for _, compra := range cliente.Compras {
		total += compra
	}
	return total
}

func calcularMediaCompras(totalCompras float64, cliente Cliente) float64 {
	if len(cliente.Compras) == 0 {
		return 0.0
	}
	return totalCompras / float64(len(cliente.Compras))
}
