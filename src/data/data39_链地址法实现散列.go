package main

import "fmt"

/**
	声明单链表
 */
type Line struct {
	val  int   //值域
	next *Line //下一个指针域
}

/**
	hashBox
 */
type hashBox struct {
	firstLine *Line //第一个指针域
	length    int   //长度
}

/**
	hashTable
 */
type HashTable2 struct {
	addr   []hashBox //hashBox切片
	length int       //元素数目
}

/**
	hash函数
 */
func hash2(key int) int {
	return key % 12
}

/**
	初始化
 */
func initHash2() HashTable2 {
	//初始化hashTable
	hashTable := new(HashTable2)

	//初始化
	for i := 0; i <= 100 ; i ++ {
		//初始化box
		box := new(hashBox)

		//添加
		hashTable.addr = append(hashTable.addr,*box)
	}

	//返回
	return *hashTable

}

/**
	添加操作
 */
 func addHash2(key int,hashTable *HashTable2) bool {
 	//获取hash
 	hash := hash2(key)

 	//构造新节点
 	newNode := new(Line)
 	newNode.val = key

 	//采用头插法
 	newNode.next = hashTable.addr[hash].firstLine
 	hashTable.addr[hash].firstLine = newNode

 	//返回
 	return  true
 }

/**
	查询操作
 */
 func findHash2 (key int,hashTable *HashTable2) bool {
 	//声明返回
 	rt := false

 	//获取hash
 	hash := hash2(key)

 	//获取head
 	head := hashTable.addr[hash].firstLine

 	//判断
 	for head != nil {
 		//判断值
 		if head.val == key {
 			rt = true
 			break
		}
 		head = head.next
	}

 	//返回
 	return rt
 }

/**
	测试
 */
func main() {
	//初始化
	hashTable := initHash2()

	//添加元素
	arr := []int{12,67,56,16,25,37,22,29,15,47,48,34}

	for _,v := range arr {
		addHash2(v,&hashTable)
	}

	fmt.Println(findHash2(49,&hashTable))
}