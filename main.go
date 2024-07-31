package main

import (
	"fmt"

	exercice "go-exercices/exercices"
)

func main() {
	fmt.Println(exercice.RepeatAlpha("abc"))
	fmt.Println(exercice.RepeatAlpha("Choumi."))
	fmt.Println(exercice.RepeatAlpha(""))
	fmt.Println(exercice.RepeatAlpha("abacadaba 01!"))
}
