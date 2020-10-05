package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxSize int
	array   [5]int //数组 =>模拟队列
	head    int    //表示队列的首
	tail    int    //表示指向队列的尾部
}

//向队列中添加元素
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列以满")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return

}

//向队列中移出元素
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列为空")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.head
	return
}

//显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下")
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//判断环形队列满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判断环形队列为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key int
	var value int
	for {
		fmt.Println("1,表示添加到队列中")
		fmt.Println("2,表示从队列中获取数据")
		fmt.Println("3,显示队列")
		fmt.Println("4,表示退出队列")
		fmt.Println("请输入[1-4]")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入你要添加到队列的值")
			fmt.Scanln(&value)
			err := queue.Push(value)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			queue.ListQueue()
		case 2:
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("队列中取出的数", val)
		}

	}
}
