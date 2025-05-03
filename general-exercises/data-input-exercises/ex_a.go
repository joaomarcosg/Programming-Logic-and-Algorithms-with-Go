package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// elaborar um programa que leia um número. Calcule e informe o dobro desse número

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe um número: ")
	scanner.Scan()

	input := scanner.Text()

	number, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Println("Informe um número inteiro")
		return
	}

	double := number * 2

	fmt.Println("Dobro é: ", double)

}
