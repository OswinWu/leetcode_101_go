package maxProfit

func maxProfit(prices []int) int {
	totalDays := len(prices)
	if totalDays < 2 {
		return 0
	} else if totalDays == 2 {
		if prices[1] > prices[0] {
			return prices[1] - prices[0]
		} else {
			return 0
		}
	}
	profit := 0
	for boughtDay := 0; boughtDay < totalDays; boughtDay++ {
		soldDay := boughtDay
		for j := boughtDay + 1; j < totalDays; j++ {
			if prices[soldDay] < prices[j] {
				soldDay = j
			} else {
				break
			}
		}
		profit = profit + prices[soldDay] - prices[boughtDay]
		boughtDay = soldDay
	}
	return profit
}
