package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//elaborar um programa que leia o peso de uma ração em kg e o quanto um gato consumo por dia, em gramas
//informe quantos dias irá durar a ração e o quanto sobra da ração(em gramas)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Peso da ração (kg): ")
	scanner.Scan()

	weight := scanner.Text()

	weigthRation, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		fmt.Println("Informe o peso em Kg")
		return
	}

	fmt.Print("Consumo diário (gr): ")
	scanner.Scan()

	consumption := scanner.Text()

	dailyConsumption, err := strconv.ParseFloat(consumption, 64)
	if err != nil {
		fmt.Println("Informe o consumo em gr")
		return
	}

	weigthRationGr := weigthRation * 1000

	duration := math.Floor(weigthRationGr / dailyConsumption)
	remaning := weigthRationGr - (dailyConsumption * duration)

	fmt.Printf("Duração: %0.f dias\nSobra: %0.fgr", duration, remaning)

}
