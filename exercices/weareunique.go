package exercice

func WeAreUnique(str1 , str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	s1 := make(map[rune]bool)
	s2 := make(map[rune]bool)
	for _, v := range str1 {
		s1[v]=true
	}
	for _, v := range str2 {
		s2[v]=true
	}
	for i := range s1 {
		for j := range s2 {
			if i==j {
				delete(s1, i)
				delete(s2, i)
			}
		}
	}
	return len(s1)+len(s2)
}
