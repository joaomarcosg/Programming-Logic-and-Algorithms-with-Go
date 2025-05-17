package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// elaborar um programa para um cinema, que leia o título e a duração de um filme em minutos.
// exiba o título do filme e converta a duração para horas e minutos.

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe o nome do filme: ")
	scanner.Scan()

	movie := scanner.Text()

	fmt.Println("Informe a duração do filme (min): ")
	scanner.Scan()

	movieDuration := scanner.Text()

	duration, err := strconv.ParseFloat(movieDuration, 64)
	if err != nil {
		fmt.Println("Informe o tempo de duração em minutos")
		return
	}

	fmt.Println(movie)
	convertToHoursAndMinutes(duration)

}

func convertToHoursAndMinutes(duration float64) {
	hours := math.Floor(duration / 60)
	minutes := math.Mod(duration, 60)

	fmt.Printf("Horas: %v\nMinutos: %v", hours, minutes)
}
