package main

import "fmt"

const Null = -100 //默认值
const Flow = -200 //表示在溢出区

/**
	声明hashTable
 */
type hashTable3 struct {
	addr []int //正常区
	flow []int //缓冲区---最好是排好序的插入可以利用插值排序等查找
	flowIndex int //缓冲区存储下标===队列的尾类似
}

/**
	hash函数
 */
func hash3(key int) int {
	return key % 12
}

/**
	初始化 == 正常区101长度
 */
 func initHash3() hashTable3 {
 	//声明hashTable
 	hash := new(hashTable3)

 	//初始化
 	for i := 0; i <= 100; i ++ {
 		hash.addr = append(hash.addr,Null)
	}

 	//返回
 	return *hash

 }

func addHash3(key int,hashTable *hashTable3) bool {
	//生成hash
	hash := hash3(key)

	//判断值
	if hashTable.addr[hash] == Null {//直接添加
		hashTable.addr[hash] = key

	} else if hashTable.addr[hash] != Flow{ //表示第一次冲突
		hashTable.flow = append(hashTable.flow,hashTable.addr[hash])
		hashTable.flow = append(hashTable.flow,key)
		hashTable.addr[hash] = Flow
	} else {
		hashTable.flow = append(hashTable.flow,key)
	}

	//返回
	return true
}
/**
	查找
 */
func findHash3(key int,hashTable *hashTable3) bool {
	//生成hash
	hash := hash3(key)

	//判断
	if hashTable.addr[hash] == Null{
		return false
	}

	//判断是否在溢出区
	if hashTable.addr[hash] == Flow {
		//查找溢出区
		n := len(hashTable.flow) - 1

		//判断
		if hashTable.flow[n] == key {
			return true
		}

		//添加哨兵
		hashTable.flow[n] = key

		//开始
		i := 0
		for hashTable.flow[i] != key {
			i++
		}

		//判断
		if i == n {
			return false
		}

		//返回
		return true

	} else {
		return hashTable.addr[hash] == key
	}

}
/**
	测试
 */
func main() {
	//初始化
	hashTable := initHash3()

	//添加元素
	arr := []int{12,67,56,16,25,37,22,29,15,47,48,34}

	for _,v := range arr {
		addHash3(v,&hashTable)
	}

	fmt.Println(findHash3(16,&hashTable))
}