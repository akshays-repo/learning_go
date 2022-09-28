package main

import (
	"fmt"
	"time"
)

func main() {

	// create channel
	ch := make(chan string)

	// function call with goroutine
	go receiveData(ch)

	fmt.Println("No data. Receive Operation Blocked")

	// send data to the channel
	ch <- "Data Received. Receive Operation Successful"

	time.Sleep(time.Second * 1)

	go receiveData(ch)

	ch <- "Data Received 2. Receive Operation Successful"

	// create channel
	number := make(chan int)
	message := make(chan string)

	// function call with goroutine
	go channelData(number, message)

	// retrieve channel data
	fmt.Println("Channel Data:", <-number, <-message)
	// fmt.Println("Channel Data:", <-message)

}

func receiveData(ch chan string) {
	fmt.Println("I am from receiveData ")
	// receive data from the channel
	fmt.Println(<-ch)

}

func channelData(number chan int, message chan string) {

	// send data into channel
	time.Sleep(time.Second * 1)
	number <- 15
	time.Sleep(time.Second * 1)
	message <- "Learning Go channel"

}
