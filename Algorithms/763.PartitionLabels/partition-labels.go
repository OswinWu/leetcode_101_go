package partitionLabels

func partitionLabels(ratings string) []int {
	splitLens := []int{}
	letterRange := map[rune][]int{}
	letterOrder := []rune{}
	for index, letter := range ratings {
		if _, found := letterRange[letter]; found {
			letterRange[letter][1] = index
		} else {
			letterOrder = append(letterOrder, letter)
			letterRange[letter] = []int{index, index}
		}
	}
	right := letterRange[letterOrder[0]][1]
	left := letterRange[letterOrder[0]][0]
	leftLen := len(ratings)
	for index, letter := range letterOrder {
		if index == 0 {
			continue
		}
		if right < letterRange[letter][0] {
			splitLen := right - left + 1
			splitLens = append(splitLens, splitLen)
			leftLen = leftLen - splitLen
			left = letterRange[letter][0]
			right = letterRange[letter][1]
		} else {
			if letterRange[letter][1] > right {
				right = letterRange[letter][1]
			}
		}
	}
	if len(splitLens) == 0 {
		splitLens = append(splitLens, len(ratings))
		leftLen = 0
	}
	if leftLen != 0 {
		splitLens = append(splitLens, leftLen)
	}
	return splitLens
}
