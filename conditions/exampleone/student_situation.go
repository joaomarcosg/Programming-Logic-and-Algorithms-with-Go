package exampleone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputData() (string, float64, float64, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o nome do aluno: ")
	studentName, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, 0, err
	}
	studentName = strings.TrimSpace(studentName)

	fmt.Print("Digite a primeira nota: ")
	firstGradeStr, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, 0, err
	}
	firstGrade, err := strconv.ParseFloat(strings.TrimSpace(firstGradeStr), 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("primeira nota inválida: %v", err)
	}

	fmt.Print("Digite a segunda nota: ")
	secondGradeStr, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, 0, err
	}
	secondGrade, err := strconv.ParseFloat(strings.TrimSpace(secondGradeStr), 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("segunda nota inválida: %v", err)
	}

	return studentName, firstGrade, secondGrade, nil

}
