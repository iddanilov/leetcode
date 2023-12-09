package main

type Value *int

type MyStack struct {
	value []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.value = append(this.value, x)
	for i := 0; i < len(this.value)-1; i++ {
		this.value = append(this.value, this.value[0])
		this.value = this.value[1:]
	}
}

func (this *MyStack) Pop() int {
	temp := this.value[0]
	this.value = this.value[1:]

	return temp
}

func (this *MyStack) Top() int {
	return this.value[0]
}

func (this *MyStack) Empty() bool {
	return len(this.value) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
