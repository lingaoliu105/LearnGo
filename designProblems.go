package main

import "math/rand"

type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (s *Solution) Reset() []int {
	return s.data
}

func (s *Solution) Shuffle() []int {
	cpy := make([]int, len(s.data))
	copy(cpy, s.data)
	rand.Shuffle(len(s.data), func(i, j int) {
		cpy[i], cpy[j] = cpy[j], cpy[i]
	})
	return cpy
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

//MinStack() 初始化堆栈对象。
//void push(int val) 将元素val推入堆栈。
//void pop() 删除堆栈顶部的元素。
//int top() 获取堆栈顶部的元素。
//int getMin() 获取堆栈中的最小元素。

type MinStack struct {
	nums []int
	min  int
}

func Constructor() MinStack {
	return MinStack{[]int{}, 0}
}

func (this *MinStack) Push(val int) {
	if len(this.nums) == 0 {
		this.min = val
	}
	if val < this.min {
		this.min = val
	}
	this.nums = append(this.nums, val)
}

func (this *MinStack) Pop() {
	popped := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	if popped == this.min && len(this.nums) > 0 {
		this.min = getMin(this.nums)
	}
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func getMin(nums []int) int {
	min := nums[0]
	for _, value := range nums {
		if value < min {
			min = value
		}
	}
	return min
}
