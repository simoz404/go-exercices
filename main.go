package main

import (
	"fmt"
	exercice "go-exercices/exercices"
	"os"
)


func main() {
	fmt.Println(exercice.Union(os.Args[1], os.Args[2]))
}