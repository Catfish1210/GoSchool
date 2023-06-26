package goschool

import "fmt"

func E004() {
	c := make(chan int)
	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	// recieve
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit")

}
