package exercice

import (
	"fmt"
	"os"
)

func Grouping(exp string, s string) {
	if exp == "" || s == "" || exp[0] != '(' || exp[len(exp)-1] != ')' {
		os.Exit(0)
	}
	indPipe := 0
	for i, v := range exp {
		if v == '|' {
			indPipe = i
		}
	}
	var exp1 string
	var exp2 string
	if indPipe == 0 {
		exp1 = exp[1 : len(exp)-1]
	} else {
		exp1 = exp[1:indPipe]
		exp2 = exp[indPipe+1 : len(exp)-1]
	}
	var s1 string
	l1 := len(exp1)
	l2 := len(exp2)
	t1, t2 := true, true
	n := 1
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '.' && s[i] != ',' {
			s1 += string(s[i])
		} else {
			for j := 0; j < len(s1); j++ {
				if j+l1 <= len(s1) && exp1 == s1[j:j+l1] && t1{
					fmt.Println(string(n+48) + ": " + s1)
					n++
					t1 = false
				}
				if (j+l2 <= len(s1) && exp2 == s1[j:j+l2]) && exp2 != "" && t2{
					fmt.Println(string(n+48) + ": " + s1)
					n++
					t2 = false
				}
			}
			s1 = ""
			t1, t2 = true, true
		}
	}
	if s1 != "" {
		fmt.Println(string(n+48) + ": " + s1)
	}
}
