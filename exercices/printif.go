package exercice

func PrintIf(str string) string {
	if len(str) > 3 || len(str) == 0 {
		return "G\n"
	}
	return "Invalid Output\n"
}
