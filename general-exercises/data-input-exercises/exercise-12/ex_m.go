package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//elaborar um programa que leia descrição e preço de um produto. Após apresente as mensagens indicando a
//promoção: 50% de desconto (para um item) na compra de três unidades do produto

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Produto: ")
	scanner.Scan()

	product := scanner.Text()

	fmt.Print("Preço R$: ")
	scanner.Scan()

	productPrice := scanner.Text()

	price, err := strconv.ParseFloat(productPrice, 64)
	if err != nil {
		fmt.Println("Informe o valor em R$")
		return
	}

	discount := price * 0.5
	totalPurchase := (2 * price) + discount

	fmt.Printf(
		"%s - Promoção: Leve 3 por R$: %2.f\nO 3° produto custa apenas R$: %2.f",
		product, totalPurchase, discount)

}
