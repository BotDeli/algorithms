package findSubstringInString

// Алгоритм поиска подстроки в строке методом Бойера-Мура-Хорспула

func findSubstringInString(str, substr string) int {
	if substr == "" {
		return 0
	}
	charSet := getCharSet(substr)
	ln, lnS := len(str), len(substr)
	cursor := lnS - 1
	moveLen := lnS
	var countSuccessful int
	for cursor < ln {
		var skip bool
		if value, have := charSet[str[cursor]]; have {
			moveLen = value
			if str[cursor] == substr[lnS-1-countSuccessful] {
				countSuccessful++
				if countSuccessful == lnS {
					return cursor
				}
				cursor--
			} else {
				skip = true
			}
		} else {
			skip = true
		}
		if skip {
			cursor += moveLen + countSuccessful
			countSuccessful = 0
			moveLen = lnS
		}
	}
	return -1
}

func getCharSet(str string) map[byte]int {
	ln := len(str)
	charSet := make(map[byte]int)
	for ind := ln - 2; ind >= 0; ind-- {
		if _, have := charSet[str[ind]]; !have {
			charSet[str[ind]] = ln - ind - 1
		}
	}
	if _, have := charSet[str[ln-1]]; !have {
		charSet[str[ln-1]] = ln
	}
	return charSet
}
