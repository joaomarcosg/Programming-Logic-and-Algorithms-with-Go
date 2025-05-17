package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// elaborar um programa para uma lan house que deve ler o valor de cada 15 minutos de uso e o tempo
// de uso por um cliente em minutos. Informe o valor a ser pago pelo cliente sabendo que as frações
// extras de 15 minutos devem ser cobradas de forma integral

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Valor por 15min de uso R$: ")
	scanner.Scan()

	valuePerTime := scanner.Text()

	value, err := strconv.ParseFloat(valuePerTime, 64)
	if err != nil {
		fmt.Println("Informe o valor em R$")
		return
	}

	fmt.Print("Tempo de uso do cliente: ")
	scanner.Scan()

	useTime := scanner.Text()

	_time, err := strconv.ParseInt(useTime, 10, 64)
	if err != nil {
		fmt.Println("Informe o tempo de uso em minutos:")
	}

	amountToPay := math.Ceil(float64(_time)/15) * value

	fmt.Printf("Valor a pagar R$: %2.f", amountToPay)

}
