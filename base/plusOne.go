func plusOne(digits []int) []int {
	len := len(digits)

	for i := len - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			for j := i + 1; j < len; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	temp := make([]int, len+1)
	temp[0] = 1
	return temp
}
