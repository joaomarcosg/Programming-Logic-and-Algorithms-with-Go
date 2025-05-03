package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o seu nome")

	scanner.Scan()

	name := scanner.Text()

	fmt.Println("Olá! ", name)
}
