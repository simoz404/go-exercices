package exercice

func Union(s1 string, s2 string) string {
	s := s1 + s2
	str := string(s[0])
	for _, v1 := range s {
		for i, v2 := range str {
			if v1 == v2 {
				break
			} else if i == len(str)-1 {
				str += string(v1)
			}
		}
	}
	return str
}
