package exercice

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	n := a[0]
	for i:=1; i < len(a); i++ {
			n = f(n, a[i])
	}
	fmt.Println(n)
}
