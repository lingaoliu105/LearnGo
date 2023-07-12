package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("finish")
		fmt.Println("inside goroutine")
		//time.Sleep(5 * time.Second)
		c <- 666 //send 666 into c
		fmt.Println("goroutine done")
	}()
	time.Sleep(5 * time.Second)
	num, ok := <-c //receive from c
	fmt.Println(num, ok)
	time.Sleep(1 + time.Millisecond)

	//buffered channel
	bc := make(chan int, 3)

	//check size
	fmt.Println(len(bc), cap(bc))

	go func() {
		defer fmt.Println("finish 2")
		for i := 0; i < 3; i++ {
			bc <- i
			fmt.Println("sub go running,lenc: ", len(bc), cap(bc))

		}
		close(bc)
	}()

	//for i := 0; i <3; i++ {
	//	num = <-bc
	//	fmt.Println(num)
	//}

	for {
		//当channel被close且再无数据，ok为false
		if data, ok := <-bc; ok { //if语句的新奇写法
			fmt.Println(data)
		} else {
			break
		}
	}

	//channel 通过 range读取,和上面代码等效
	for data := range bc {
		fmt.Println(data)
	}
	var x int
	//channel and select：对多路channel进行监控
	for {
		select {
		case c <- x: //c 可写时
			continue
		case <-c: //c 可读时
			continue
		}
	}
}

//注：channel不需要经常关闭（unlike files）
//关闭channel后无法再写入数据但仍可获取剩余数据
//nil channel收发都会被阻塞
