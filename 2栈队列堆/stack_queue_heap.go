package main

import "github.com/golang-collections/collections/queue"

// 1.用队列实现栈
type MyStack struct {
	data *queue.Queue
}

//
func NewMyStack() *MyStack {
	return &MyStack{
		data: queue.New(),
	}
}

// 元素入栈
func (ms *MyStack) push(value interface{}) {
	tmpQueue := queue.New()
	// 元素入临时队列
	tmpQueue.Enqueue(value)

	for !ms.empty() {
		// 先将栈内部队列中元素入临时队列
		tmpQueue.Enqueue(ms.pop())
	}

	// 将临时队列中元素全部入栈的内部队列
	for tmpQueue.Len() != 0 {
		ms.data.Enqueue(tmpQueue.Dequeue())
	}
}

// 弹出栈顶元素
func (ms *MyStack) pop() interface{} {
	return ms.data.Dequeue()
}

// 栈顶元素
func (ms *MyStack) top() interface{} {
	return ms.data.Peek()
}

// 栈为空
func (ms *MyStack) empty() bool {
	return ms.data.Len() == 0
}
