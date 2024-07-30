package exercice

func RevConcatAlternate(slice1,slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return []int{}
	}
	max := len(slice1)-1
	if len(slice2)>len(slice1) {
		max = len(slice2)-1
	}
	var s []int
	for i := max; i >= 0; i-- {
		if i < len(slice1) {
			s = append(s, slice1[i])
		}
		if i < len(slice2) {
			s = append(s, slice2[i])
		}
	}
	return s
}
