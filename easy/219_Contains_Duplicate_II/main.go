package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))

}

func containsNearbyDuplicate(nums []int, k int) bool {
	temp := make(map[int]int)
	for numb, value := range nums {
		if val, ok := temp[value]; ok {
			if numb-val <= k {
				return true
			}
		}
		temp[value] = numb
	}

	return false
}
