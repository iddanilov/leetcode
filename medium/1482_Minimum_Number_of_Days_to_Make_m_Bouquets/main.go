package main

import "fmt"

func main() {
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1))
}

func minDays(bloomDay []int, mm int, k int) int {
	if (mm * k) > len(bloomDay) {
		return -1
	}
	max := 0
	r := 0

	for i := range bloomDay {
		if bloomDay[i] > max {
			max = bloomDay[i]
		}
	}

	for max > r {
		m := (max + r) / 2
		vv := 0
		c := 0
		for i := range bloomDay {
			if bloomDay[i] <= m {
				vv++
			} else {
				if vv >= k {
					c += vv / k
				}
				vv = 0
			}
		}
		if vv >= k {
			c += vv / k
		}
		if c >= mm {
			max = m
		} else if c < mm {
			r = m + 1
		}
	}

	return max
}
