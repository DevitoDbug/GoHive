package main

import "fmt"

func FindingCommonElements(arry1 []int, arry2 []int) (commonElements []int) {

	largestArray := len(arry1)
	if len(arry2) > largestArray {
		largestArray = len(arry2)
	}

	commonArray := make([]int, 0, largestArray)

	for _, num1 := range arry1 {
		for _, num2 := range arry2 {
			if num2 == num1 {
				commonArray = append(commonArray, num1)
			} else {
				continue
			}
		}
	}

	return commonArray
}

func MoreEfficientMethod(arry1 []int, arry2 []int) (commonElements []int) {
	newArry := append(arry1, arry2...)

	seenNumber := make(map[int]bool)
	commonArry := make([]int, 0, len(newArry))

	for _, num := range newArry {
		if !seenNumber[num] {
			seenNumber[num] = true
		} else {
			commonArry = append(commonArry, num)
		}
	}
	return commonArry
}

func main() {
	arry1 := []int{2, 3, 6, 7, 9}
	arry2 := []int{9, 7}

	fmt.Println(MoreEfficientMethod(arry2, arry1))
}
