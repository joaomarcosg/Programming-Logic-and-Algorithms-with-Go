package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// elaborar um programa que leia descrição e preço de um medicamento. Informe o valor do produto na
// promoção (na compra de duas unidades do medicamento é descontado os centavos do valor total)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Medicamento: ")
	scanner.Scan()

	drug := scanner.Text()

	fmt.Print("Preço R$: ")
	scanner.Scan()

	price := scanner.Text()

	drugPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Informe o valor do medicamento em R$")
		return
	}

	promotionPrice := math.Floor(drugPrice) * 2

	fmt.Printf("Promoção de: %s\nLeve 2 por apenas R$: %.0f", drug, promotionPrice)

}
