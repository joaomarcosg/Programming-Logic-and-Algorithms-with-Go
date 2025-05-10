package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// elaborar um programa para uma pizzaria, o qual leia o valor total de uma conta e
// quantos clientes vão pagá-la. Calcule e informe o valor a ser pago por cliente.

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe o valor da conta R$: ")
	scanner.Scan()

	accountValue := scanner.Text()

	account, err := strconv.ParseInt(accountValue, 10, 64)
	if err != nil {
		fmt.Println("Informe um número")
		return
	}

	fmt.Println("Informe o número de pessoas que vão pagar: ")
	scanner.Scan()

	people := scanner.Text()
	clients, err := strconv.ParseInt(people, 10, 64)
	if err != nil {
		fmt.Println("Informe um número")
		return
	}

	fmt.Println("O valor a ser pago por cliente será R$: ", amountPerClient(account, clients))

}

func amountPerClient(account, client int64) int64 {
	var amount int64
	if account >= client {
		amount = account / client
	}

	return amount
}
