package exercice
func ZipString(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		n := 1
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			for j := i+1; j < len(s); j++ {
				if s[i] == s[j] {
					n++
				} else {
					break
				}
			}
			str += string((n+48)) + string(s[i])
			i += n-1
		} else {
			str += string((n+48))+string(s[i])
		}
	}
	return str
}
