package main

import "fmt"

type Node struct {
	no       int
	name     string
	nickname string
	prev     *Node
	next     *Node
}

//给链表插入一个节点
// func InsertNode(head *Node, newNode *Node) {
// 	temp := head
// 	for {
// 		if temp.next == nil {
// 			break //表示找到最后
// 		}
// 		temp = temp.next
// 	}
// 	temp.next = newNode
// 	newNode.prev = temp
// }

//双向链表顺序添加
func InsertNode(head *Node) {

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
	ShowNode(head)

}
