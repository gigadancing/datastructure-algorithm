package main

import (
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

// 1.用队列实现栈
type MyStack struct {
	data *queue.Queue
}

// 构造函数
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

// 2.用栈实现队列
type MyQueue struct {
	data *stack.Stack
}

// 构造函数
func NewMyQueue() *MyQueue {
	return &MyQueue{
		data: stack.New(),
	}
}

// 队列为空
func (mq *MyQueue) empty() bool {
	return mq.data.Len() == 0
}

// 元素入队列
func (mq *MyQueue) enqueue(value interface{}) {
	// 临时栈
	tmpStack := stack.New()

	// 将队列内部栈内的元素入临时栈
	for !mq.empty() {
		tmpStack.Push(mq.dequeue())
	}
	// 将元素入临时栈
	tmpStack.Push(value)
	// 将临时栈内元素重新入队列的内部栈
	for tmpStack.Len() != 0 {
		mq.data.Push(tmpStack.Pop())
	}
}

// 元素出队列
func (mq *MyQueue) dequeue() interface{} {
	return mq.data.Pop()
}

// 队列头元素
func (mq *MyQueue) peek() interface{} {
	return mq.data.Peek()
}

// 3.设计一个栈，栈的操作push(x)、pop()、top()、getMin()算法复杂度为O(1)
type MyStackWithMin struct {
	data, min *stack.Stack // min是存放最小值的栈
}

// 构造函数
func NewMyStackWithMin() *MyStackWithMin {
	return &MyStackWithMin{
		data: stack.New(),
		min:  stack.New(),
	}
}

// 栈为空
func (mswm *MyStackWithMin) empty() bool {
	return mswm.data.Len() == 0
}

// 元素入栈
func (mswm *MyStackWithMin) push(value interface{}) {
	// 元素入栈
	mswm.data.Push(value)

	if mswm.min.Len() == 0 { // 最小值栈为空
		mswm.min.Push(value)
	} else {
		if value.(int) < mswm.min.Peek().(int) { // 压入元素比最小值小
			mswm.min.Push(value)
		} else {
			mswm.min.Push(mswm.min.Peek())
		}
	}
}

// 返回栈最小元素
func (mswm *MyStackWithMin) getMin() interface{} {
	return mswm.min.Peek()
}

// 弹出栈顶元素
func (mswm *MyStackWithMin) pop() interface{} {
	mswm.min.Pop()
	return mswm.data.Pop()
}

// 返回栈顶元素
func (mswm *MyStackWithMin) top() interface{} {
	return mswm.data.Peek()
}
