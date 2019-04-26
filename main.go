package goarea

import "math"

//Pi é uma proporção numerica definida pela relação entre
//o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

//Circ é responsavel por calcular a area da circuferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//React é responsavel por calcular a área de um retangulo
func React(base, altura float64) float64 {
	return base * altura
}

//função não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
