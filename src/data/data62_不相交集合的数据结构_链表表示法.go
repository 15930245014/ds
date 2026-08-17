package main

import "fmt"

/**
	用于不相交集合的数据结构:维护了一个不相交集合的动态集合& = {S1,S2,....Sk}。我们用一个代表来表示每一个集合（一般为集合的第一个元素）
	不相交集合的常见操作:
	make-set(x):建立一个新的集合，它的唯一集合代表是x
	union(x,y):将包含x,y的2个集合合并成一个新的集合
	find-set(x):返回指向x的唯一代表的集合
 */

 /**
 	定义单链表结构体
  */
  type linkTable struct {
  		val byte
  		pre *linkTable
  		next *linkTable
  }

/**
	定义每个集合结构体
 */
 type structure struct {
 		head,tail *linkTable
 		hash map[byte]int
 }

 /**
 	定义集合
  */
  type structureArray struct {
  		arr []structure
  		hash  map[byte]int
  }
/**
	建立一个新的集合，它的代表为x
 */
func makeStructure(s structureArray,x byte) structureArray{
	//判断hash中是否存在
	if s.hash[x] == 1 {
		fmt.Println("已经存在！")
		return s
	}

	//构造新的链表元素
	newLink := new(linkTable)
	newLink.val = x

	//构造新的S
	structure := new(structure)
	structure.head = newLink
	structure.tail = newLink
	structure.hash = map[byte]int{}
	structure.hash[x] = 1

	//新链表指向head
	newLink.pre = structure.head

	//写入s
	s.arr = append(s.arr,*structure)
	s.hash[x] = 1

	//返回信息
	return s
}
/**
	查找
 */
func findStructure(s structureArray,x byte) *structure {
	//判断hash中是否存在
	if s.hash[x] == 0 {
		fmt.Println("不存在！")
		return nil
	}

	//查找
	var rtStruecure *structure
	for i := 0; i < len(s.arr);i ++ {
		if s.arr[i].head.val == x {
			rtStruecure = &(s.arr[i])
			break
		}
	}

	//返回
	return rtStruecure
}

/**
	合并体系
	加权合并的启发式策略-----将链表的小的合并到稍大的
 */
func unionStructure(x,y byte) {
	//1、查询包含x,和 y 的体系
	//2、遍历其中之一体系，另一体系的tail尾插法填入遍历体系，并修改体系的head
	//3、删除合并后的体系
}


func main() {
	//声明空的
	s := structureArray{}
	s.hash = map[byte]int{}

	s = makeStructure(s,'a')
	s = makeStructure(s,'b')
	fmt.Println(findStructure(s,'b').hash)
}

