package exercice

import "fmt"

func PrintMemory(arr [10]byte) {
	dec := "0123456789abcdef"
	var out string
	for i := 0; i < len(arr); i++ {
		div := arr[i]/16
		mod := arr[i]%16
		s := string(dec[div])+string(dec[mod])
		out += s + " "
	}
	fmt.Print(out)
}
