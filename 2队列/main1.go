package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	array   [5]int //数组 =>模拟队列
	front   int    //表示队列的首
	rear    int    //表示指向队列的尾部
}

//添加数据到队列中
func (this *Queue) AddQueue(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("队列以满")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

//显示队列
func (this *Queue) ShowQueue() {
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

//从队列中取数据
func (this *Queue) GetQueue() (val int, err error) {
	if this.rear == this.front {
		return 0, errors.New("当前队列为空")
	}
	this.front++
	val = this.array[this.front]
	return
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(value)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			queue.ShowQueue()
		case 2:
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("队列中取出的数", val)
		}

	}
}
