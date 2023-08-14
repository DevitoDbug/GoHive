package main

import "fmt"

func fib(num int) (nthNum int) {
	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func main() {
	job := make(chan int, 100)
	result := make(chan int, 100)

	//increasing the number of go routines does not infinitely increase the processing speed
	//the processing speed increases till all cores are used
	go worker(job, result)
	go worker(job, result)
	go worker(job, result)
	go worker(job, result)

	//sending values to the channel
	for i := 0; i <= 100; i++ {
		job <- i
	}
	close(job)

	//receiving values from the channel
	for i := 0; i <= 100; i++ {
		fmt.Println(<-result)
	}

}
