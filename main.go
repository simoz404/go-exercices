package main

import (
	"fmt"
	"os"

	exercice "go-exercices/exercices"
)

func main() {
	var s []string
	for i, v := range os.Args {
		if i != 0 {
			s = append(s, v)
		}
	}
	fmt.Println(exercice.Reversestrcap(s))
}
