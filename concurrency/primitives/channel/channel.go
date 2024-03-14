package main

import "fmt"


func main () {
	
	// Joint Point implementation
	myChannel := make(chan string)

	go func () {
		myChannel <- "data"
	}()

	/* 
		This is blocking line of code
		This means that our main function is actually going
		to wait for the data to be put onto this channel to receive
		it here.
	*/
	result := <-myChannel


	fmt.Println(result)
}