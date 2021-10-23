package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// using wait group
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		read("Learning about go routine")
		waitGroup.Done()
	}()

	go func() {
		read("exemple")
		waitGroup.Done()
	}()

	waitGroup.Wait()

	// using Channels

	ch := make(chan string)
	go readToChannel("test 1!", ch)
	go readToChannel("test 2!", ch)
	go readToChannel2("test 3!", ch)
/*
	for {
		message, open := <- ch
		fmt.Println(message)
		if !open {
			break
		}
	}
*/
	for message := range ch {
		fmt.Println(message)
	}
	fmt.Println("go routine executed sucessfully")

	// channels using buffer
	// channels with buffer doesnt block the program execution
	ch2 := make(chan string, 3)

	ch2 <- "Send 1 to channel"
	ch2 <- "Send 2 to channel"
	ch2 <- "Send 3 to channel"

	msg1:= <- ch2
	msg2 := <- ch2
	msg3 := <- ch2

	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
}

func read(text string) {
	for i:=0; i < 7; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func readToChannel(text string, ch chan string) {
	for i:=0; i < 7; i++ {
		ch <- text
		time.Sleep(time.Second)
	}
}

func readToChannel2(text string, ch chan string) {
	for i:=0; i < 7; i++ {
		ch <- text
		time.Sleep(time.Second * 2)
	}
	close(ch)
}