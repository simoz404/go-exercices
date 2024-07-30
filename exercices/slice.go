package exercice

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 1 {
		return a[nbrs[0]:]
	}
	return a[nbrs[0]:nbrs[1]]
}
