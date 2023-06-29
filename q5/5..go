package main

import (
	"errors"
)

func converterTemperatura(temperatura float64, escalaOrigem, escalaDestino string) (float64, error) {
	switch escalaOrigem {
	case "C":
		switch escalaDestino {
		case "F":
			return temperatura*9/5 + 32, nil
		case "K":
			return temperatura + 273.15, nil
		default:
			return 0, errors.New("escala inválida")
		}
	case "F":
		switch escalaDestino {
		case "C":
			return (temperatura - 32) * 5 / 9, nil
		case "K":
			return (temperatura-32)*5/9 + 273.15, nil
		default:
			return 0, errors.New("escala inválida")
		}
	case "K":
		switch escalaDestino {
		case "C":
			return temperatura - 273.15, nil
		case "F":
			return (temperatura-273.15)*9/5 + 32, nil
		default:
			return 0, errors.New("escala inválida")
		}
	default:
		return 0, errors.New("escala inválida")
	}
}
