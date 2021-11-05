package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := readString("Test 1", "Test 2")

	for i := 0; i < 7; i++ {
		fmt.Println(<-channel1)
		fmt.Println(<-channel2)
	}

}

// Standard Generator
func readString(text1 string, text2 string) (<-chan string, <-chan string) {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			ch1 <- fmt.Sprintf("Param1 text is: %s", text1)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			ch2 <- fmt.Sprintf("Param2 text is: %s", text2)
			time.Sleep(time.Second)
		}
	}()

	return ch1, ch2
}