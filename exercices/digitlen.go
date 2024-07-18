package exercice

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	} else if n < 0 {
		n = -n
	}
	var i int
	for n > 0 {
		n /= base
		i++
	}
	return i
}
