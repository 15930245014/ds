package main

import "fmt"

/**
	哈希表定义
 */
const NULL = -100 //默认值--
const MOD = 12    //hash函数mod值
const L = 100     //地址长度
type HashTable1 struct {
	addr   []int //保存地址
	length int   //长度
}

/**
	初始化
 */
func initHash1() HashTable1 {
	//初始化
	hashTable := new(HashTable1)

	//初始化地址
	for i := 0; i <= L; i ++ {
		hashTable.addr = append(hashTable.addr, NULL)
	}

	//返回
	return *hashTable
}

/**
	哈希函数
 */
func hash1(key int) int {
	return key % MOD
}

/**
	添加元素---开放定址法解决冲突
 */
func addHash1(key int, HashTable *HashTable1) bool {
	//获取hash
	hash := hash1(key)

	//解决冲突
	for HashTable.addr[hash] != NULL {
		hash = hash1(hash + 1)
	}

	//写入
	HashTable.addr[hash] = key

	//返回
	return true
}

/**
	查找记录
 */
func findHash1(key int, HashTable *HashTable1) bool {
	//声明返回
	rt := true

	//根据关键子获取hash地址
	hash := hash1(key)
	origin := hash

	//判断
	for HashTable.addr[hash] != key {
		//开放寻址法
		hash = hash1(hash + 1)

		//判断
		if hash == origin || HashTable.addr[hash] == NULL {
			rt = false
			break
		}

	}

	//返回
	return rt

}
/**
	测试
 */
func main() {
	//初始化
	hashTable := initHash1()

	//添加元素
	arr := []int{12,67,56,16,25,37,22,29,15,47,48,34}

	for _,v := range arr {
		addHash1(v,&hashTable)
	}

	fmt.Println(findHash1(0,&hashTable))
}