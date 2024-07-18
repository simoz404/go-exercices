package main

import (
	"fmt"

	exercice "go-exercices/exercices"
)

func main() {
	fmt.Println(exercice.DigitLen(100, 10))
	fmt.Println(exercice.DigitLen(100, 2))
	fmt.Println(exercice.DigitLen(-100, 16))
	fmt.Println(exercice.DigitLen(100, -1))
}
