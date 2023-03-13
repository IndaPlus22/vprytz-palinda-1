package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	// sum the numbers in a
	var sum int
	for _, v := range a {
		sum += v
	}

	// send the result sum to channel res
	res <- sum
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	// receive the results from the two goroutines
	x, y := <-ch, <-ch

	// return the sum of the two results
	return x + y
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
