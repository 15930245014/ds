package main

import "fmt"

/**
构建单链表的数据结构
 */
type LTable struct {
	data int  //值域
	next *LTable //指向下一个节点的指针
}

func main() {
	//测试单链表
	arr := []int{1,2,3,4}

	//声明头
	L := new(LTable)

	//头插法构造
	createHead(L,arr)

	//打印信息
	read(L)

	//尾插法
	createTail(L,arr)
	read(L)

	//利用尾插法结果为例

	//添加元素
	insert(L,2,0)
	read(L)

	//删除元素
	deleteLtable(L,2)
	read(L)


}

//读取单链表的信息

/**
*@params L 头指针
 */
func read(L *LTable) {
	p := L.next//第一个数据域
	for p != nil {
		fmt.Print((*p).data)
		p = p.next
	}

}


//头插法构造单链表
func createHead(L *LTable,arr []int) {
	//添加计数
	count := 0

	for _,value := range arr {
		tempLtable := new(LTable)

		(*tempLtable).data = value

		(*tempLtable).next = (*L).next
		(*L).next = tempLtable
		count ++
	}

	//头节点一般存放链表的长度
	(*L).data = count
}

//尾插法构造单链表 ***
func createTail(L *LTable,arr []int) {
	//构造起始节点和计数器
	start := new(LTable)
	(*start).data = arr[0]
	L.next = start

	count := 1

	for i := 1; i < len(arr);i++ {
		node := new(LTable)
		(*node).data = arr[i]
		start.next = node
		start = start.next
		count ++
	}

	(*L).data = count

}

//插入i位置插入一个元素 1成功0失败 i从1开始
func insert(L *LTable,i,key int) int {
	//定义一个结果量
	var result int

	//判断是否越界
	if (*L).data < i {
		result = 0
		return result
	}

	//寻找位置
	p := (*L).next
	//计数器
	j := 1 //由于p是从头指针往下开始的

	for p.next != nil && j < i - 1 {
		p = p.next
		i ++
	}

	//此时 p 为插入前的指针
	newNode := new(LTable)

	(*newNode).data = key

	(*newNode).next = p.next
	p.next = newNode

	(*L).data += 1

	result = 1
	return result

}
//删除元素 i从1开始
func deleteLtable(L *LTable,i int) int {
	//判断位置
	if (*L).data < i {
		return 0
	}

	p := (*L).next
	j := 1
	//寻找删除前元素
	for p != nil && j < i{
		p = p.next
		j ++
	}

	//执行删除操作
	p.next = p.next.next

	//回收内存
	(*L).data -= 1

	return 1


}

//修改元素 由于特别简单忽略
func modifyITable(L *LTable,i,key int) {

}

//删除lianbiao
func truncate(L *LTable) int {
	p := L.next

	for p != nil {
		Q := p.next
		//删除p
		p = Q
	}

	return 1
}


