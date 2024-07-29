package exercice

func Union(s1 string, s2 string) string {
	s := s1 + s2
	str := ""
	ss := make(map[rune]bool)
	for _, v1 := range s {
		if !ss[v1] {
			str += string(v1)
			ss[v1] = true
		}


	}
	return str
}
