package main

import "fmt"

func main() {
	fmt.Println(getRow(10))
}

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		tmp := row[i-1]
		tmp = tmp * (rowIndex + 1 - i) / i
		row[i] = tmp
	}
	return row
}
