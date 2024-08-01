package exercice

import (
	"fmt"
	"strconv"
)

func Romannumbers(s string) {
	n, _ := strconv.Atoi(s)
	valuesint := []int{1000, 500, 100, 50, 10, 5, 1}
	valuesrune := []rune{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	var str string
	for n > 0 {
		for i, v := range valuesint {
			if n-v >= 0 {
				str +=  string(valuesrune[i])
				n -= v
				break
			}
		}
	}
	var l string
	for i := 0; i < len(str); i++ {
		if i+4 <= len(str) && str[i:i+4] == "IIII" {
			l += "IV"
			i+= 3
		} else if i+4 <= len(str) && str[i:i+4] == "XXXX" {
			l += "XL"
			i+=3
		} else if i+4 <= len(str) && str[i:i+4] == "CCCC" {
			l += "CD"
			i+=3
		} else {
			l += string(str[i])
		}
	}

	fmt.Println(l)
}
