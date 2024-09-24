package exam4

func Ft_coin(coins []int, amount int) int {
	a := 0
	b := 0
	if amount == 0 {
		return 0
	} else if amount < 0 {
		return -1
	}
	for i := len(coins) - 1; i >= 0; i-- {
		for j := 0; a < amount; j++ {
			a += coins[i]
			b += 1
		}
		if a > amount {
			a -= coins[i]
			b -= 1
		}
	}
	if a != amount {
		return -1
	}
	return b
}
