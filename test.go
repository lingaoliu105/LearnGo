package main

import "fmt"

func main() {
	sl := []int{1, 2, 3}
	ptr := (*[3]int)(sl)
	fmt.Println(ptr)
	sl2 := append(sl, 4)
	fmt.Println(sl2)
	ptr2 := (*[4]int)(sl2)
	fmt.Println(ptr2)
}
