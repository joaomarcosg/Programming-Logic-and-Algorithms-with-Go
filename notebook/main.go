package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Println(_RandInt())
	fmt.Println(_RandFloat())
	fmt.Println(_Sqrt(16))

}

// retorna o valor absoluto de um número
func _Abs(num float64) float64 {
	return math.Abs(num)
}

// arredonda o valor para cima
func _Ceil(num float64) float64 {
	return math.Ceil(num)
}

// arredonda o valor para baixo
func _Floor(num float64) float64 {
	return math.Floor(num)
}

// retorna a base elevada ao expoente
func _Pow(base, exp float64) float64 {
	return math.Pow(base, exp)
}

// retorna um inteiro aleatório
func _RandInt() int {
	return rand.Intn(10)
}

// retorna um ponto flutuante aleatório
func _RandFloat() float64 {
	return rand.Float64()
}

// retorna a raiz quadrada do número
func _Sqrt(num float64) float64 {
	return math.Sqrt(num)
}
