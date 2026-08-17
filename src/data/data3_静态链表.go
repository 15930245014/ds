package main

import "fmt"

const MAXSIZE = 10  //定义静态链表的值域最大长度

/**
	静态链表
   总体分为空闲链表区和占有链表区
  数组的0 cur指向空闲链表的index 数组的最后一项cur指向值域开始
  cur = 0 代表空 类似单链表的nil 一般是值域最后一个元素
 */
type staticTable struct {
	data int   //利用数组来保存值域
	cur int  //指向下一个索引的序号
}

func main() {
	//利用切片来创建静态链表
	arr := []int{1,2,3}

	//创建静态链表
	sTArray := creteStaticTable(arr)

	//插入一个数值4
	sTArray  = addStaticTable(sTArray,2,4)

	readStaticTable(sTArray)

	sTArray = deleteStaticTable(sTArray,2)
	readStaticTable(sTArray)

	fmt.Print(sTArray)


}

//遍历静态链表
func readStaticTable(sTArray [MAXSIZE]staticTable) {
	//查看值域指向
	index := sTArray[MAXSIZE - 1].cur

	for sTArray[index].cur != 0 {
		fmt.Println(sTArray[index].data)
		index = sTArray[index].cur
	}

	//读取最后一个元素
	fmt.Println(sTArray[index].data)
}

/**
创建静态链表
 */
func creteStaticTable(arr []int) [MAXSIZE]staticTable {
	//声明一个静态链表数组
	sTArray := [MAXSIZE]staticTable{}

	//由于顺序的从1开始添加
	for i,value := range arr {
		//声明
		tempNode := staticTable{value,i + 2} //cur 指向下一个元素 因为存放为i+1 它要指向下一个故为 i+1+1

		sTArray[i + 1] = tempNode   //由于指向空闲索引 故从1开始

	}

	//最后一个记录占有区开始
	sTArray[MAXSIZE - 1] = staticTable{0,1}

	//第一个记录空闲区索引
	sTArray[0] = staticTable{0,len(arr) + 1}

	//空闲区初始化cur
	for k := len(arr) + 1; k < MAXSIZE -1;k++ {
		sTArray[k] = staticTable{0, k + 1}
	}

	//最后一个节点cur指向0
	sTArray[len(arr)].cur = 0

	//返回静态链表数组
	return sTArray

}

//添加静态链表 i从1开始 在i之后添加key

func addStaticTable(sTArray [MAXSIZE]staticTable,i,key int)[MAXSIZE]staticTable {
	//先放入空闲区 然后在改变指向
	index := sTArray[0].cur //空闲区索引占有重新指向新的
	sTArray[0].cur = sTArray[index].cur

	//空闲节点存放值
	sTArray[index] = staticTable{key,0}

	//插入操作
	sTArray[index].cur = sTArray[i].cur
	sTArray[i].cur = index

	return sTArray

}

//删除静态链表 i从1开始
func deleteStaticTable(sTArray [MAXSIZE]staticTable,i int) [MAXSIZE]staticTable {
	//先寻找元素之前的节点
	//查看值域指向
	index := sTArray[MAXSIZE - 1].cur
	j := 1
	for sTArray[index].cur != 0 && j < i - 1 {

		index = sTArray[index].cur
	}

	//获取删除的元素的index
	deleteCur := sTArray[index].cur

	//删除元素
	sTArray[index].cur = sTArray[deleteCur].cur

	//空闲区修改
	sTArray[deleteCur].cur = sTArray[0].cur
	sTArray[0].cur = deleteCur

	return  sTArray

}

//修改略
func modifyStaticTable() {

}

