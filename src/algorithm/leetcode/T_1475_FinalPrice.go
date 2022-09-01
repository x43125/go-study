package main

import "fmt"

func main() {
	prices := []int{8, 4, 6, 2, 3}
	res := finalPrices(prices)
	for _, re := range res {
		fmt.Print(re, " ")
	}
}

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))

	for i, price := range prices {
		res[i] = price
		for j := i + 1; j < len(prices); j++ {
			if price >= prices[j] {
				res[i] = price - prices[j]
				break
			}
		}
	}
	return res
}
