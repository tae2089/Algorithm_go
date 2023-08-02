package leetcode

func isValid(s string) bool {
	// create stack
	stack := make([]byte, 0)

	//create map for string
	var closingToOpening = map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != closingToOpening[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
