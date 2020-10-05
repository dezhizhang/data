### 单链表
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
### 有顺序的单链表
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



