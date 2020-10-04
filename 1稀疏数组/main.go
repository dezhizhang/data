package main

import "fmt"

type Node struct {
	row int
	col int
	val int
}

func main() {

	//1,创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//存盘退出
	var sparseArr []Node
	for index, item := range chessMap {
		for number, value := range item {
			if value != 0 {
				node := Node{
					row: index,
					col: number,
					val: value,
				}
				sparseArr = append(sparseArr, node)
			}
		}
	}
	//续上盘
	var chessMap2 [11][11]int
	for _, value := range sparseArr {
		chessMap2[value.row][value.col] = value.val
	}

	//恢复后的数据
	for _, item := range chessMap2 {
		for _, val := range item {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}
