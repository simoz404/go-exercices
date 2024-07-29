package exercice

func ThirdTimeIsACharm(str string) string {
	var s string
	for i, v := range str {
		if i%3 == 2 {
			s += string(v)
		}
	}
	return s+"\n"
}
