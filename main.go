package main

import (
	"os"

	exercice "go-exercices/exercices"
)

func main() {
	// arrByte := [2048]byte{}
	// fmt.Println(arrByte[1]-1)
	exercice.Brainfuck(os.Args[1])
}
