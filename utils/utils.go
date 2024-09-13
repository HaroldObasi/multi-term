package utils

func SplitRuneArray(runes []rune, sep rune) [][]rune {
	var result [][]rune
	var temp []rune

	for _, r := range runes {
		if r == sep {
			if len(temp) > 0 {
				result = append(result, temp)
				temp = []rune{} // Reset temp
			}
		} else {
			temp = append(temp, r)
		}
	}

	// Append the last part if not empty
	if len(temp) > 0 {
		result = append(result, temp)
	}

	return result
}
