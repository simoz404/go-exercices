package main

import (
	"os"

	exercice "go-exercices/exercices"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	exercice.Grouping(os.Args[1], os.Args[2])
}
