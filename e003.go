package goschool

import "fmt"

func E003() {
	c := make(chan int)

	go foo(c)

	bar(c)
	fmt.Println("about to exit")

}

//send only channel
func foo(c chan<- int) {
	c <- 42
}

//recieve only channel
func bar(c <-chan int) {
	fmt.Println(<-c)
}
