package main

import (
	"fmt"

	exercice "go-exercices/exercices"
)

func main() {
	fmt.Println(exercice.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(exercice.RetainFirstHalf("A"))
	fmt.Println(exercice.RetainFirstHalf(""))
	fmt.Println(exercice.RetainFirstHalf("Hello World"))
}
