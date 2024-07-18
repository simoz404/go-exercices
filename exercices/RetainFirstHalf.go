package exercice

func RetainFirstHalf(str string) string {
	var i int
	if len(str) > 1 {
		i = len(str)/2
	} else {
		i = len(str)
	}
	return str[:i]
}
