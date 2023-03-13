package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	// indefinite loop
	for {
		// print time and text
		now := time.Now()
		fmt.Println("The time is " + now.Format("15.04.05") + ": " + text)

		// sleep for delay
		time.Sleep(delay)
	}
}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	go Remind("Time to sleep", 60*time.Second)

	// wait
	select {}
}
