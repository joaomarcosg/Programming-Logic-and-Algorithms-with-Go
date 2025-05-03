package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Elaborar um programa que leia um número. Calcule e informe os seus vizinhos, ou seja, o número
	// anterior e posterior.

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe um número: ")

	scanner.Scan()

	input := scanner.Text()

	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Informe um número inteiro")
		return
	}

	fmt.Println("Vizinhos: ", number-1, ",", number+1)

}
