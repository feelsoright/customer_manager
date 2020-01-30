package utils

import "fmt"

func PrintLine(content string) {
	fmt.Printf("----------------------------%v----------------------------\n", content)
}

func YOrN(content string) bool {
	for {
		val := GetStringColumn(content)
		if val == "y" {
			return true
		}

		if val == "n" {
			return false
		}
		fmt.Println("您的输入有误，请重新输入！")
	}
}
