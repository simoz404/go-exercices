package main

import (
	"fmt"

	exercice "go-exercices/exercices"
)


func main() {
	fmt.Print(exercice.PrintIfNot("abcdefz"))
	fmt.Print(exercice.PrintIfNot("abc"))
	fmt.Print(exercice.PrintIfNot(""))
	fmt.Print(exercice.PrintIfNot("14"))
}
