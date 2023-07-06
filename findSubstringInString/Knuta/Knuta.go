package Knuta

// Алгоритм Кнута-Морриса-Пратта (КМП-алгоритм)

func findSubstringInString(str, substr string) int {
	if substr == "" {
		return 0
	}
	ln, lnS := len(str), len(substr)
	slicePrefixLen := getArrayLengthPrefix(substr)
	var strCur, substrCur int
	for strCur < ln {
		if str[strCur] == substr[substrCur] {
			strCur++
			substrCur++
			if substrCur == lnS {
				return strCur - lnS
			}
		} else {
			substrCur = 0
			strCur += slicePrefixLen[substrCur] + 1
		}
	}
	return -1
}

func getArrayLengthPrefix(str string) []int {
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
