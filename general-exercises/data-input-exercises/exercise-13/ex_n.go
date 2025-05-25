package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//elaborar um programa que leia o salário e o tempo que um funcionário trabalha na empresa.
//Sabendo que a cada 4 anos (quadriênio) o funcionário recebe um acrescimo de 1% no salário,
//calcule e informe o número de quadriênios a que o funcionário tem direito e o salário final

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Salário R$: ")
	scanner.Scan()

	employeeSalary := scanner.Text()

	salary, err := strconv.ParseFloat(employeeSalary, 64)
	if err != nil {
		fmt.Println("Informe o salário em R$: ")
		return
	}

	fmt.Println("Tempo (anos): ")
	scanner.Scan()

	workedTime := scanner.Text()

	time, err := strconv.ParseFloat(workedTime, 64)
	if err != nil {
		fmt.Println("Informe o tempo em anos")
		return
	}

	quadriennium := math.Floor(time / 4)
	addition := salary * quadriennium / 100
	finalSalary := salary + addition

	fmt.Printf("Quadriênios: %0.f\nSalário final R$: %2.f", quadriennium, finalSalary)

}
