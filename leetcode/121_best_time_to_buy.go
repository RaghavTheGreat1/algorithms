package main

func MaxProfit(prices []int) int {
	minIndex, minValue := 0, prices[0]
	profit := 0

	for i, v := range prices {
		if v < minValue {
			minIndex, minValue = i, v
		}

		if i > minIndex {
			tempProfit := v - minValue
			if tempProfit > profit {
				profit = tempProfit
			}
		}
	}

	if profit < 0 {
		return 0
	}

	return profit
}
