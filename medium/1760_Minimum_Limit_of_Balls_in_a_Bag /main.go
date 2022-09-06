package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSize([]int{7, 17}, 2))
}

func minimumSize(nums []int, maxOperations int) int {
	l := 1
	max := 0
	a := int(^uint(0) >> 1)

	// find max value
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}

	for l <= max {
		m := (l + max) / 2
		c := cnt(m, nums)
		if c <= maxOperations {
			if a > m {
				a = m
			}
		}
		if c > maxOperations {
			l = m + 1
		} else {
			max = m - 1
		}
	}

	return a

}

func cnt(v int, nums []int) int {
	c := 0
	for _, i := range nums {
		if i <= v {
			continue
		}
		x := i / v
		if x%1 == 0 {
			c += x - 1
		} else {
			c += x
		}

	}
	return c
}
