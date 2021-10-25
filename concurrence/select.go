package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			ch1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 3)
			ch2 <- "Channel 2"
		}
	}()

	for {
		select {
		case msgChannel1 := <- ch1:
			fmt.Println(msgChannel1)

		case msgChannel2 := <- ch2:
			fmt.Println(msgChannel2)
		}
	}
}
