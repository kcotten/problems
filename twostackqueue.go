package main

type MyQueue struct {
	front  int
	s1, s2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: []int{},
		s2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	if len(this.s1) == 0 {
		this.s1 = append(this.s1, x)
		this.front = x
	} else {
		// copy to s2 and append
		for len(this.s1) > 0 {
			this.s2 = append(this.s2, this.s1[0])
			this.s1 = this.s1[1:]
		}
		this.s2 = append(this.s2, x)
		// copy back to s1 for next op
		for len(this.s2) > 0 {
			this.s1 = append(this.s1, this.s2[0])
			this.s2 = this.s2[1:]
		}
	}
}

func (this *MyQueue) Pop() int {
	x := this.s1[0]
	this.s1 = this.s1[1:]
	if len(this.s1) > 0 {
		this.front = this.s1[0]
	} else {
		this.front = 0
	}
	return x
}

func (this *MyQueue) Peek() int {
	return this.front
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}
