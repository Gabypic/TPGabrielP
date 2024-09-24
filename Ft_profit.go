package exam4

func Ft_profit(prices []int) int {
	a := prices[0]
	b := 0
	date := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < a {
			a = prices[i]
			date = i
		}
	}
	for j := date; j < len(prices); j++ {
		if prices[j] > b {
			b = prices[j]
		}
	}
	return b - a
}
