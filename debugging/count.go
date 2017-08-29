package main

import "time"

func count(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}

func main() {
	println("Starting main")
	c := make(chan int)
	go count(c)
	for i := range c {
		println("count:", i)
	}
}
