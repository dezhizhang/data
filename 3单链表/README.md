# 单链表
### 定义 
单链表）是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要通过顺序读取从头部开始；链表是使用指针进行构造的列表；又称为结点列表，因为链表是由一个个结点组装起来的；其中每个结点都有指针成员变量指向列表中的下一个结点；
列表是由结点构成，head指针指向第一个成为表头结点，而终止于最后一个指向nuLL的指针。
### 优点
单个结点创建非常方便，普通的线性内存通常在创建的时候就需要设定数据的大小
结点的删除非常方便，不需要像线性结构那样移动剩下的数据
结点的访问方便，可以通过循环或者递归的方法访问到任意数据，但是平均的访问效率低于线性表。
### 单链表图
![单链表](https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=3380107229,75248244&fm=26&gp=0.jpg)
### 代码实现
1，无序添加
```
package main

import "fmt"

type Node struct {
	no       int
	name     string
	nickname string
	next     *Node
}

//给链表插入一个节点
func InsertNode(head *Node, newNode *Node) {
	//定义一个临时变量
	temp := head
	for {
		if temp.next == nil { //表示找到最后一个节点
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	//加入到链表的最后
	temp.next = newNode
}

func ShowNode(head *Node) {
	temp := head
	if temp.next == nil {
		fmt.Println("当前链表为空")
		return
	}

	for {
		fmt.Println("[%d,%s,%d] ==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	head := &Node{}
	node1 := &Node{
		no:       1,
		name:     "宋江",
		nickname: "及时兩",
	}
	node2 := &Node{
		no:       2,
		name:     "鲁智深",
		nickname: "智多星",
	}
	InsertNode(head, node1)
	InsertNode(head, node2)
	ShowNode(head)

}

```
2，有顺序的单链表
```
package main

import "fmt"

type Node struct {
	no       int
	name     string
	nickname string
	next     *Node
}

//给链表插入一个节点
func InsertNode(head *Node, newNode *Node) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			break
		} else if temp.next.no == newNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("当前节点以存在", newNode.no)
		return
	} else {
		newNode.next = temp.next
		temp.next = newNode
	}
}

//删除一个节点
func DelNode(head *Node, no int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == no {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("没有找到你要删除的节点")
	}

}

func ShowNode(head *Node) {
	temp := head
	if temp.next == nil {
		fmt.Println("当前链表为空")
		return
	}

	for {
		fmt.Printf("[%d,%s,%s] ==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	head := &Node{}
	node1 := &Node{
		no:       1,
		name:     "宋江",
		nickname: "及时兩",
	}
	node2 := &Node{
		no:       2,
		name:     "鲁智深",
		nickname: "智多星",
	}
	node3 := &Node{
		no:       3,
		name:     "刘德华",
		nickname: "华子",
	}
	InsertNode(head, node1)
	InsertNode(head, node3)
	InsertNode(head, node2)
	DelNode(head, 2)
	ShowNode(head)

}
```



