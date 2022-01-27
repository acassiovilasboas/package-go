package area

import "math"

// Pi é uma constante
const Pi = 3.1416

// Circulo retorna a area de um circulo
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Retangulo retorna a area de um retangulo
func Retangulo(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}