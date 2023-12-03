package main

func convertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}
	columnNumber--
	charCode := 'A' + rune(columnNumber%26)
	return convertToTitle(columnNumber/26) + string(charCode)
}
