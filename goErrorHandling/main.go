package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, err := strconv.Atoi("234th")
	if err != nil {
		fmt.Printf("Can't convert: %q", err)
		num = 0
	}
	fmt.Print(num)
}
