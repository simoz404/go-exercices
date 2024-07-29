package exercice

func Slice(a []string, nbrs ...int) []string {
	l1 := len(a)
	l2 := len(nbrs)
	var i, j int
	if l2 == 1 {
		if nbrs[0] < 0 {
			if l1+nbrs[0] > l1 {
				return []string{}
			}
			return a[l1+nbrs[0]:]
		} else {
			if nbrs[0] > l1 {
				return []string{}
			}
			return a[nbrs[0]:]
		}
	} else {
		if nbrs[0] < 0 && nbrs[1] < 0 {
			if l1+nbrs[0] > l1+nbrs[1] {
				return nil
			}
			if l1+nbrs[0] > l1 || l1+nbrs[1] < 0 {
				return []string{}
			}
			i = l1+nbrs[0]
			j = l1+nbrs[1]
			if l1+nbrs[1] > l1 {
				j = l1
			}
			if l1+nbrs[0] < 0 {
				i = 0
			}
			return a[i:j]
		} else if nbrs[0] > 0 && nbrs[1] < 0 {
			if l1+nbrs[0] > nbrs[1] {
				return nil
			}
			if nbrs[0] > l1 || l1+nbrs[1] < 0 {
				return []string{}
			}
			i = nbrs[0]
			j = l1+nbrs[1]
			if l1+nbrs[1] > l1 {
				j = l1
			}
			if nbrs[0] < 0 {
				i = 0
			}
			return a[i:j]
		} else if nbrs[0] < 0 && nbrs[1] > 0 {
			if l1+nbrs[0] > nbrs[1] {
				return []string{}
			}
			if l1+nbrs[0] > l1 || nbrs[1] < 0 {
				return []string{}
			}
			i = l1+nbrs[0]
			j = nbrs[1]
			if nbrs[1] > l1 {
				j = l1
			}
			if l1+nbrs[0] < 0 {
				i = 0
			}
			return a[i:j]
		} else {
			if nbrs[0] > nbrs[1] {
				return nil
			}
			if nbrs[0] > l1  {
				return []string{}
			}
			i = nbrs[0]
			j = nbrs[1]
			if nbrs[1] > l1 {
				j = l1
			}
			if nbrs[0] < 0 {
				i = 0
			}
			return a[i:j]
		}
	}
}
