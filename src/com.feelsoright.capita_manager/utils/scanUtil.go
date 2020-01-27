package utils

import "fmt"

func GetStringColumn(content string) string {
	column := ""
	fmt.Printf("%v:", content)
	_, _ = fmt.Scanln(&column)
	return column
}

func GetIntColumn(content string) int {
	column := 0
	fmt.Printf("%v:", content)
	_, _ = fmt.Scanln(&column)
	return column
}
