package exercice

import "fmt"

func Brainfuck(s string) {
	arrByte := [2048]byte{}
	pointer := 0
	indArr := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			arrByte[pointer]++
		} else if s[i] == '-' {
			arrByte[pointer]--
		} else if s[i] == '>' {
			pointer++
		} else if s[i] == '<' {
			pointer--
		} else if s[i] == '[' {
			indArr = append(indArr, i)
		} else if s[i] == ']' {
			if arrByte[pointer] != 0 {
				i = indArr[len(indArr)-1]
			} else {
				indArr = indArr[:len(indArr)-1]
			}
		} else if s[i] == '.' {
			fmt.Print(string(arrByte[pointer]))
		}
	}
	fmt.Println()
}
