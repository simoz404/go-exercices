package exercice

func Reversestrcap(str []string) string {
	var s string
	var s1 string
	for _, v := range str {
		for _, v1 := range v {
			if v1 >= 'A' && v1 <= 'Z' {
				s1 += string(v1 + 32)
			} else {
				s1 += string(v1)
			}
		}
		for i := 0; i < len(s1); i++ {
			if i+1 < len(s1) && (s1[i+1] < 'a' || s1[i+1] > 'z') && s1[i+1] != '\'' {
				if s1[i] >= 'a' && s1[i] <= 'z' {
					s += string(s1[i]-32)
				} else {
					s += string(s1[i])
				}
			} else if i == len(s1)-1 && (s1[i] >= 'a' && s1[i] <= 'z'){
				s += string(s1[i]-32)
			} else {
				s += string(s1[i])
			}
		}
		s1 = ""
		s += "\n"
	}
	return s
}
