package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(2 << 2)
	fmt.Println(reorderedPowerOf2(13))
}

func reorderedPowerOf2(N int) bool {
	signatureN := makeSignature(N)
	for i := 0; i < 32; i++ {
		if makeSignature(1<<i) == signatureN {
			return true
		}
	}
	return false
}

func makeSignature(n int) int {
	if n == 0 {
		return 0
	}
	leading, remaining := n/10, n%10
	return makeSignature(leading) + int(math.Pow(10, float64(remaining)))
}
