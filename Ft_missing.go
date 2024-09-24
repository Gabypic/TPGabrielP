package exam4

func Ft_missing(nums []int) int {
	var test bool
	for i := 0; i < len(nums); i++ {
		test = false
		for j := 0; j < len(nums); j++ {
			if nums[j] == i {
				test = true
			}
		}
		if test == false {
			return i
		}
	}
	orderNums := tri(nums)
	return orderNums[len(orderNums)-1] + 1
}

func tri(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
