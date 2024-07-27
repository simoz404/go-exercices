package main

import (
	"fmt"
	exercice "go-exercices/exercices"
	"os"
)


func main() {
	if len(os.Args) != 3 {
		return
	}
	fmt.Println(exercice.Wdmatch(os.Args[1], os.Args[2]))
}
