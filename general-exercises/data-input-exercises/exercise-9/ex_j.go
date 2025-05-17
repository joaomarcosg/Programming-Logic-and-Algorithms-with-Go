package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// elaborar um programa para um restaurante que leia o preço por kg e o consumo (em gramas)
// exiba o valor a ser pago.

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Buffet por Quilo R$: ")
	scanner.Scan()

	buffetPerKilo := scanner.Text()

	valuePerKilo, err := strconv.ParseFloat(buffetPerKilo, 64)
	if err != nil {
		fmt.Println("Insira o valor em R$")
		return
	}

	fmt.Print("Consumo do cliente (gr): ")
	scanner.Scan()

	constumeConsumption := scanner.Text()

	consumption, err := strconv.ParseFloat(constumeConsumption, 64)
	if err != nil {
		fmt.Println("Insira o consumo em gramas(gr)")
	}

	amountToPay := valuePerKilo * (consumption / 1000)

	fmt.Printf("Valor a ser pago R$: %.2f", amountToPay)

}
