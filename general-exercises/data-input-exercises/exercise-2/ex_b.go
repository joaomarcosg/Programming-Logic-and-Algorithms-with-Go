package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Elaborar um programa que leia dois números. Calcule e informe a soma desses números.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe um número: ")
	scanner.Scan()

	firstNumber := scanner.Text()

	num1, err := strconv.ParseInt(firstNumber, 10, 64)
	if err != nil {
		fmt.Println("Informe um número inteiro")
		return
	}

	fmt.Println("Informe outro número: ")
	scanner.Scan()

	secondNumber := scanner.Text()

	num2, err := strconv.ParseInt(secondNumber, 10, 64)
	if err != nil {
		fmt.Println("Informe um número inteiro")
		return
	}

	sum := num1 + num2

	fmt.Println("Soma é", sum)

}
