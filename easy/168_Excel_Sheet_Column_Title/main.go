package main

import "fmt"

func main() {
	fmt.Println(convertToTitle2(123))

}

func convertToTitle2(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}
	columnNumber--
	charCode := 'A' + rune(columnNumber%26)
	fmt.Println(string(charCode))
	fmt.Println(convertToTitle2(columnNumber / 26))
	return convertToTitle2(columnNumber/26) + string(charCode)
}
