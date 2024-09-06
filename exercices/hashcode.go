package exercice

func HashCode(dec string) string {
	l := len(dec)
	var s string
	var n int
	for _, v := range dec {
		n = (int(v)+l)%127
		if n >= 0 && n <= 127 {
			s += string(n)
		} else {
			s += string(n+33)
		}
	}
	return s
}
