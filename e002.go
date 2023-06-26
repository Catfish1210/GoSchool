package goschool

import "fmt"

func E002() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
