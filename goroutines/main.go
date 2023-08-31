package main

import (
	"fmt"
	"time"
)

func printMsg(text string, channel chan string) {
	for i := 0; i < 3; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)

	}

	channel <- "CHANN"

}

func main() { //main goroutine
	channel := make(chan string)
	go printMsg("Hello world", channel)
	// go printMsg("Hello again")
	// go printMsg("Hello again and again")

	response := <-channel
	fmt.Print(response)

}
