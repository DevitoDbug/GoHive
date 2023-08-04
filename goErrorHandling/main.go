package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, err := strconv.Atoi("234th")
	if err != nil {
		fmt.Printf("lkajsdlfkjaldsjflakjdslfkajsd", err)
		num = 0
	}
	fmt.Print(num)
}
