package main

import "fmt"

func hammingWeight(num uint32) int {
	i := 0
	for num > 0 {
		i = i + int(num&1)
		num = num >> 1

	}
	return i
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
