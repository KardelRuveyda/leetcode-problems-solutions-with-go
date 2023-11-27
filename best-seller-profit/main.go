package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	res := 0

	for i := range prices {
		debug()
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else if prices[i]-minPrice > res {
			res = prices[i] - minPrice
		}
	}

	return res
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(nums)
	fmt.Println(result)
}
