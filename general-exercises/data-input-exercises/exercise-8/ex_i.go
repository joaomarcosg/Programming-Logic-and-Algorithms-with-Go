package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// elaborar um programa para uma revenda de veículos. O programa deve ler modelo e preço do veículo.
// apresentar como resposta o valor da entrada (50%) e o saldo em 12x.

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Veículo: ")
	scanner.Scan()

	vehicleModel := scanner.Text()

	fmt.Print("Preço R$: ")
	scanner.Scan()

	vehiclePrice := scanner.Text()

	price, err := strconv.ParseFloat(vehiclePrice, 64)
	if err != nil {
		fmt.Println("Informe o valor do veículo em R$")
		return
	}

	downPayment := price * 0.50
	installmentPayment := (price - downPayment) / 12

	fmt.Printf("Promoção: %s\n", vehicleModel)
	fmt.Printf("Entrada de R$: %.2f e +12x de R$: %.2f\n", downPayment, installmentPayment)

}
