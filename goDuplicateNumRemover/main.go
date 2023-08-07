package main

import (
	"fmt"
)

func RemoveDuplicate(inputs []int) (output []int) {

	uniqueNums := make([]int, 0, len(inputs))
	seenValue := make(map[int]bool)

	for _, num := range inputs {
		if !seenValue[num] {
			uniqueNums = append(uniqueNums, num)
			seenValue[num] = true
		}
	}
	return uniqueNums
}

func main() {
	testNum := []int{1, 2, 2, 2, 3, 3, 3, 3, 4, 5, 6}
	result := RemoveDuplicate(testNum)
	fmt.Println(result)
}
