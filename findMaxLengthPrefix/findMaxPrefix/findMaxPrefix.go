package findMaxPrefix

func findMaxLengthPrefix(str string) int {
	if len(str) < 2 {
		return 0
	}
	lengths := make([]int, len(str))
	left, right := 0, 1
	var max int
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
		if max < lengths[right-1] {
			max = lengths[right-1]
		}
	}
	return max
}
