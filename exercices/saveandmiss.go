package exercice

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	var s string
	n := num
	for i := 0; i < len(arg); i++ {
		if i < num {
			s += string(arg[i])
		}
		if i == num {
			i += n-1
			num += n+n
		}
	}
	return s
}
