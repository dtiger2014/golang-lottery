package main

import (
	"fmt"
	"time"
)
// 无buffer的channel，如果没有接受另外一段接受channel数据，就会阻塞
// 在等待不会到来的它（数据）。自己判断没发继续，就自杀（deadlock）了。

func worker(id int, c chan int) {
	// for {
	// 	fmt.Printf("worker %d received %c\n", id, <-c)
	// }

	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}


// channel 可以return
// chan<- int : 返回，只能输入的channel
// <-chan int : 返回，只能读取的channel
// chan int : 返回，能读取/输入的channel
func createWorker(id int) chan<- int {
	// 1. 创建channel 可以使用make(chan int)
	channel := make(chan int)

	go worker(id, channel)

	return channel
}

func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Second)
}

func main() {
	// chanDemo()
	bufferedChannel()
}

// ========== buffer channel
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}