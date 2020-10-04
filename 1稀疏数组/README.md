# 稀疏数组
### 什么是稀疏数组
稀疏数组可以看做是普通数组的压缩，但是这里说的普通数组是值无效数据量远大于有效数据量的数组
```
              0 0 0 0 0 0 0 0 0 0 0
              0 0 1 0 0 0 0 0 0 0 0
              0 0 0 0 2 0 0 0 0 0 0
              0 0 0 0 0 0 0 0 0 0 0
              0 0 0 0 0 0 0 0 0 0 0
              0 0 0 0 0 0 0 0 0 0 0
              0 0 0 0 0 0 0 0 0 0 0
              0 0 0 0 0 0 0 0 0 0 0
              0 0 0 0 0 0 0 0 0 0 0
              0 0 0 0 0 0 0 0 0 0 0
              0 0 0 0 0 0 0 0 0 0 0
```
### 其稀疏数组形式
```
              11 11 2
              1  2  1
              2  4  2
```
### 存储
1,刚说到稀疏数组是一种压缩后的数组，为什么要进行压缩存储呢？  
2,原数组中存在大量的无效数据，占据了大量的存储空间，真正有用的数据却少之又少  
3,压缩存储可以节省存储空间以避免资源的不必要的浪费，在数据序列化到磁盘时，压缩存储可以提高IO效率  
### 存储方式 
1,普通存储 
![普通存储](https://img-blog.csdnimg.cn/20190703122416582.png)
2,链式存储 
![链式存储](https://img-blog.csdnimg.cn/20190703122823779.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MTkyMjI4OQ==,size_16,color_FFFFFF,t_70)
3,十字链式存储 
![十字链式存储](https://img-blog.csdnimg.cn/20190703123032956.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MTkyMjI4OQ==,size_16,color_FFFFFF,t_70)
### 代码实现
```
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

```
### 结果
```
0       0       0       0       0       0       0       0       0       0       0
0       0       1       0       0       0       0       0       0       0       0
0       0       0       2       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
0       0       0       0       0       0       0       0       0       0       0
```








