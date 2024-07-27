package exercice

func Wdmatch(s1 string, s2 string) string {
	var n int
	var m int
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				if j > m {
					n++
					m = j
					break
				}
			}
		}
	}
	if n != len(s1) {
		return ""
	}
	return s1
}
