package main

import (
	"fmt"
	"os"

	exercice "go-exercices/exercices"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(exercice.Brackets(os.Args[1]))
	}
}
