package main

import "fmt"

func main() {
	str := "hello, 123"
	n := 0
	_, err := fmt.Sscanf(str, "hello, %d", &n)
	if err != nil {
		return
	}
}
