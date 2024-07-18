package main

import (
	"fmt"

	exercice "go-exercices/exercices"
)

func main() {
	fmt.Print(exercice.PrintIf("abcdefz"))
	fmt.Print(exercice.PrintIf("abc"))
	fmt.Print(exercice.PrintIf(""))
	fmt.Print(exercice.PrintIf("14"))
}
