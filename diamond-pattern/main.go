package main

import "fmt"

func main() {
	star := "* "
	num := 3
	for i := 0; i < num; i++ {
		for j := 0; j < num-i; j++ {

			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {

			fmt.Print(star)
		}
		fmt.Print("\n")
	}
	for i := 0; i < num; i++ {
		for j := 0; j < i; j++ {

			fmt.Print(" ")
		}

		for j := 0; j < num-i; j++ {

			fmt.Print(star)
		}
		fmt.Print("\n")
	}
}
