package main

import (
	"fmt"
	"time"
)

func foo(num string) {
	fmt.Println(num)
}

func main() {

	go foo("1")
	go foo("2")
	go foo("3")

	time.Sleep(time.Second * 2)

	fmt.Println("hi")
}