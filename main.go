package main

import (
	"fmt"

	exercice "go-exercices/exercices"
)

func main() {
	fmt.Println(exercice.SaveAndMiss("123456789", 3))
	fmt.Println(exercice.SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(exercice.SaveAndMiss("", 3))
	fmt.Println(exercice.SaveAndMiss("hello you all ! ", 0))
	fmt.Println(exercice.SaveAndMiss("what is your name?", 0))
	fmt.Println(exercice.SaveAndMiss("go Exercise Save and Miss", -5))
}
