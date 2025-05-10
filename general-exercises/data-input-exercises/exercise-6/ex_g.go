package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// elaborar um programa para uma loja, o qual leia o preço de um produto e informe as opções de pagamento
// da loja. Calcule e informe o valor para pagamento a vista com 10% de desconto e o valor em 3x.

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe o valor do produto R$: ")
	scanner.Scan()

	productValue := scanner.Text()

	value, err := strconv.ParseFloat(productValue, 64)
	if err != nil {
		fmt.Println("Informe um número: ")
		return
	}

	discountedPrice := value - (value * 0.10)
	installmentValue := value / 3

	fmt.Printf("Valor à vista com desconto R$: %2.f\nOu em 3x de R$: %2.f", discountedPrice, installmentValue)

}
