package main

import "fmt"

func writer(channel chan int, word string) {
	for true {
		<-channel
		fmt.Println(word)
		channel <- 0
	}
}

func main() {
	firstChannel := make(chan int)
	go func() { writer(firstChannel, "Hello") }()
	go func() { writer(firstChannel, "World") }()
	go func() { writer(firstChannel, "Welcome") }()
	firstChannel <- 0
	for true {

	}
}
