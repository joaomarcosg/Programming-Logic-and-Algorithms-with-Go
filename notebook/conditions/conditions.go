package conditions

import (
	"fmt"
	"math"
)

const color = "Blue"
const year = 2017

func CheckOrConditions() {

	// parênteses indica ordem de precedência das comparações
	if (color == "Blue" || color == "White") && year == 2017 {
		fmt.Println("primeiro compara Blue or White")
	}

	if color == "Blue" && (year == 2016 || year == 2017) {
		fmt.Println("primeiro compara year 2016 or year 2017")
	}

}

func ShortStatementCondition() {

	// cria variáveis nas condicionais if
	if x := math.Sqrt(4); x < 10 {
		fmt.Println(x)
	}

}
