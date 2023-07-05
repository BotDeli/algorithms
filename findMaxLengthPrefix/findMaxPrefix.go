package findMaxLengthPrefix

func findMaxLengthPrefix(str string) int {
	if len(str) < 2 {
		return 0
	}
	lengths := getArrayMaxLengthPrefix(str)
	return lengths[len(str)-1]
}

func getArrayMaxLengthPrefix(str string) []int {
	lengths := make([]int, len(str))
	left, right := 0, 1
	for right < len(str) {
		if str[left] == str[right] {
			lengths[right] = left + 1
			right++
			left++
		} else if left == 0 {
			lengths[right] = 0
			right++
		} else {
			left = lengths[left-1]
		}
	}
	return lengths
}
