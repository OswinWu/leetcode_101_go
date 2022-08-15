package candy

func candy(ratings []int) int {
	size := len(ratings)
	childrenCandies := make([]int, size)
	index := 1
	for {
		if index == size {
			break
		}
		if ratings[index] > ratings[index-1] {
			childrenCandies[index] = childrenCandies[index-1] + 1
		}
		index++
	}

	index = size - 1
	for {
		if index == 0 {
			break
		}
		if (ratings[index-1] > ratings[index]) && (childrenCandies[index-1] <= childrenCandies[index]) {
			childrenCandies[index-1] = childrenCandies[index] + 1
		}
		index--
	}
	allCandies := size
	for _, v := range childrenCandies {
		allCandies = allCandies + v
	}
	return allCandies
}
