package canPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	maxPlant := 0
	maxLen := len(flowerbed)
	if maxLen == 1 {
		if flowerbed[0] == 0 {
			maxPlant = maxPlant + 1
		}
	} else if maxLen == 2 {
		if (flowerbed[0] + flowerbed[1]) == 0 {
			maxPlant = maxPlant + 1
		}
	} else if maxLen >= 3 {
		for i := 0; i < maxLen; i++ {
			if i == 0 {
				if (flowerbed[i] + flowerbed[i+1]) == 0 {
					maxPlant = maxPlant + 1
					flowerbed[0] = 1
				}
			} else if (i + 1) == maxLen {
				if (flowerbed[i] + flowerbed[i-1]) == 0 {
					maxPlant = maxPlant + 1
				}
			} else {
				if (flowerbed[i-1] + flowerbed[i] + flowerbed[i+1]) == 0 {
					flowerbed[i] = 1
					maxPlant = maxPlant + 1
				}
			}
		}
	}
	if maxPlant >= n {
		return true
	} else {
		return false
	}
}
