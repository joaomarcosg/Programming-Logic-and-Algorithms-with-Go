package main

import "fmt"

func main() {

	studentName, firstGrade, secondGrade, err := InputData()
	if err != nil {
		fmt.Println("Erro na entrada de dados", err)
		return
	}

	average := (firstGrade + secondGrade) / 2

	if average >= 7 {
		fmt.Printf("Parabéns %s! Você foi aprovado(a)\n", studentName)
		return
	}

	fmt.Printf("Ops %s... você foi reprovado(a)\n", studentName)

}
