package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new goroutine", i)
		time.Sleep(1 * time.Millisecond)
	}
}
func main() {

	//创建goroutine，执行newTask()
	//go newTask()

	//i := 0
	//for {
	//	i++
	//	fmt.Println("main goroutine: ", i)
	//	time.Sleep(1 * time.Millisecond)
	//}

	//父goroutine结束时，子也会结束

	//通过匿名函数进行goroutine
	go func() bool {
		defer fmt.Println("a.defer")
		func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")

			//在内部函数中退出goroutine：(在外层函数的话直接return即可)
			runtime.Goexit()
		}() //define and invoke anonymous func
		fmt.Println("A")
		return true //外界无法通过直接赋值拿到该返回值
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
