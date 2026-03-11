package main

import "fmt"

func CheckStatus(s string) {
	if s == "new" {
		fmt.Println("new")
		return
	}

	if s == "old" {
		fmt.Println("old")
		return
	}
}

func main() {
	CheckStatus("new")
}
