package main

import "fmt"


func sliceToChannel(nums []int) <-chan int {
	res := make(chan int)
	
	go func () {
		for _, n := range nums {
			res <- n
		}
		close(res)
	}()
	return res
}

func square(in <-chan int) <-chan int {
	res := make(chan int)
	
	go func () {
		for n := range in {
			res <- n * n
		}
		close(res)
	}()
	return res
}


func main() {
	// input

	nums := []int{5 , 3, 4 ,7}
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	squaredChannel := square(dataChannel)
	// stage 3
	for n := range squaredChannel{
		fmt.Println(n)
	}
}