package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := multiplexer(readString("learning GO"), readString("Testing multiplexer"))

	for i := 0; i < 7; i++ {
		fmt.Println(<-ch)
	}
}

// Standard Multiplexer
func multiplexer(ch1, ch2 <-chan string) <-chan string {
	leaveChannel := make(chan string)

	go func() {
		for {
			select {
			case msg := <-ch1:
				leaveChannel <- msg
			case msg := <-ch2:
				leaveChannel <- msg
			}
		}
	}()
	return leaveChannel
}

func readString(text1 string) <-chan string {
	ch1 := make(chan string)

	go func() {
		for {
			ch1 <- fmt.Sprintf("Param1 text is: %s", text1)
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		}
	}()

	return ch1
}
