package main

import (
	"fmt"

	exercice "go-exercices/exercices"
)

func main() {
	fmt.Println(exercice.CamelToSnakeCase("HelloWorld"))
	fmt.Println(exercice.CamelToSnakeCase("helloWorld"))
	fmt.Println(exercice.CamelToSnakeCase("camelCase"))
	fmt.Println(exercice.CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(exercice.CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(exercice.CamelToSnakeCase("hey2"))
}
