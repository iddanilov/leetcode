package main

import "fmt"

var lol = []int{0, 1, 0, 3, 12}

func main() {
	moveZeroes(lol)
}
func moveZeroes(nums []int) {
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			fmt.Println("before: ", nums)
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			fmt.Println("after: ", nums)
			nonZeroIndex++
		}
	}
}
