package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Elaborar um programa que leia o valor de um jantar. Calcule e informe o valor da taxa do garçom (10%)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe o valor do jantar R$: ")
	scanner.Scan()
	dinnerPrice := scanner.Text()

	value, err := strconv.ParseFloat(dinnerPrice, 64)
	if err != nil {
		fmt.Println("Informe um número!")
		return
	}

	waiter := value * 0.10
	total := value + waiter

	fmt.Printf("Taxa do garçom R$: %2.f\nTotal R$: %2.f", waiter, total)

}
