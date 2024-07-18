package exercice

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	var str string
	for i, v := range s {
		if isAlphaBig(byte(v)) && isAlphaBig(s[i+1]) {
			return s
		} else if i-1 >= 0 && i+1 < len(s) && isAlphaSmall(s[i-1]) && isAlphaSmall(s[i+1]) &&  isAlphaBig(byte(v)) {
			str += "_" + string(v)
		} else {
			str += string(v)
		}
	}
	return str
}

func isAlphaSmall(r byte) bool {
	return r >= 'a' && r <= 'z'
}

func isAlphaBig(r byte) bool {
	return r >= 'A' && r <= 'Z'
}