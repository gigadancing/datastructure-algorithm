package main

import (
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

// 1.用队列实现栈
// 思路：用一个临时队列调换顺序
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
// 思路：用一个临时栈调换顺序
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
// 思路：用一个栈存放每次压栈、出栈后的最小值
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

// 4.已知1至n的数字序列，按顺序入栈，每个数字入栈后即可出栈，也可在栈中停留，
// 等待后面的数字入栈出栈后，该数字即可出栈入栈，求该数字序列的出栈顺序是否合法。
// 思路：用一个队列存放出栈顺序，将元素安顺序据入栈。每入一个元素，检查栈顶和队列头元素是否相等，若相等，弹出栈顶元素，弹出队列头元素；
// 否则，继续压栈；最后，若栈为空，则说明顺序合法。
func CheckOrder(order *queue.Queue) bool {
	n := order.Len()
	if n <= 1 {
		return true
	}
	s := stack.New()
	for i := 1; i <= n; i++ {
		// 按顺序入栈
		s.Push(i)
		// 栈顶元素和队列头元素相等
		for s.Len() != 0 && order.Peek() == s.Peek() {
			// 分别弹出栈顶元素和队列头元素
			s.Pop()
			order.Dequeue()
		}
	}

	if s.Len() != 0 {
		return false
	}

	return true
}

// 5.设计一个计算器，输入一个字符串存储的数学表达式，包含“(”，“)”，“+”，“-”四中符号的数学表达式
// 输入的数学表达式字符串保证是合法的，其中可能含有空格字符。
func compute(numStack, optStack *stack.Stack) {
	if numStack.Len() < 2 {
		return
	}
	// 取数据栈最上面的两个数
	num2 := numStack.Peek().(int)
	numStack.Pop()
	num1 := numStack.Peek().(int)
	numStack.Pop()
	// 取操作符
	if optStack.Peek() == '+' {
		numStack.Push(num1 + num2)
	} else if optStack.Peek() == '-' {
		numStack.Push(num1 - num2)
	}
	optStack.Pop()
}
