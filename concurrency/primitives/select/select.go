package main

import "fmt"

func main () {
	
	// Joint Point implementation
	myChannel := make(chan string)
	anotherChannel := make(chan string)
	
	go func () {
		myChannel <- "data"
	}()

	go func () {
		anotherChannel <- "data to another channel"
	}()
	
	select {
	case msgFromMyChannel := <- myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromMyAnotherChannel := <- anotherChannel:
		fmt.Println(msgFromMyAnotherChannel)

	}

}