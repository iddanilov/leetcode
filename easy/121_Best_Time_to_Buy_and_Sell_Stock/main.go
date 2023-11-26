package main

import "fmt"

func maxProfit(prices []int) int {
	temp, min := 0, prices[0]
	for _, price := range prices {
		if min > price {
			min = price
		} else {
			if temp < price-min {
				temp = price - min
			}
		}
	}
	return temp
}

func main() {
	fmt.Println(maxProfit([]int{7, 3, 7, 1, 3, 3}))
}
