package exercice

func RepeatAlpha(s string) string {
	var str string
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			for l := 0; l <= int(v-97); l++ {
				str += string(v)
			}
		} else {
			str+= string(v)
		}
	}
	return str
}
