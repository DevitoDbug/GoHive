package main

import "fmt"

func main() {
	// decimal := 3.0
	fname := "David"
	sname := "Ochieng"

	// fmt.Printf("The varrible is of type %T", decimal)

	fullname := fmt.Sprintf("My full names are %s %s ", fname, sname)
	fmt.Print(fullname)
}
