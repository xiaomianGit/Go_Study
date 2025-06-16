func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			fmt.Println("a")
			stack = append(stack, char)
		case ')', '}', ']':
			fmt.Println(stack)
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if pairs[char] != top {
				return false
			}
			len := len(stack) - 1
			stack = stack[:len]
		}
	}

	fmt.Println(stack)
	return len(stack) == 0
}
