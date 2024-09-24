package exam4

func Ft_max_substring(s string) int {
	b := []rune(s)
	a := 1
	c := []int{}
	for i := 0; i < len(b)-1; i++ {
		if b[i+1] == b[i]+1 {
			a++
		}
		if b[i] == b[0] && i > 0 {
			c = append(c, a)
		}
	}
	if The_Higher(c) > 1 {
		return The_Higher(c) - 1
	}
	return 1
}

func The_Higher(tab []int) int {
	a := 0
	for i := 0; i < len(tab); i++ {
		if tab[i] > a {
			a = tab[i]
		}
	}
	return a
}
