package main

import "fmt"

/**
循环链表：循环链表是在单链表的基础上最后节点的next指向头结点 但访问最后一个元素十分困难 故引入了尾指针 指向最后一个节点
tail->next 即为头节点 tail->next->next 为第一个元素
 */

 /**
 循环链表的数据结构和单链表相同
  */
  type xhTable struct {
  		data int
  		next *xhTable
  }

  func main() {
  	//利用循环链表构建循环链表
  	arr := []int{1,2,3}

  	//构建头指针
  	L := new(xhTable)

  	//构建循环链表
  	tail := createXhTable(L,arr)

  	//遍历循环链表
  	readXhTable(tail)

  	//循环链表的插入 、删除和单链表不同之处 h = *tail.next 为头指针  略

  	//测试合并循环链表
  	L2 := new(xhTable)
  	tail2 := createXhTable(L2,arr)
  	tail2 = mergeXhTable(tail,tail2)
  	readXhTable(tail2)



  }



  /**
  利用尾插法构建循环链表
  返回尾指针 ===本例为最后一个元素
   */
  func createXhTable(L *xhTable,arr []int) *xhTable {
  	  //实力化第一个元素
  	  start := new(xhTable)

	  (*start).data = arr[0]
	  (*start).next = nil //默认为nil

	  //头指针指向第一个元素
	  (*L).next = start

	  //开始构建遍历
	  for i := 1; i < len(arr);i++ {
	  	  //声明临时
	  	  temp := new(xhTable)
		  (*temp).data = arr[i]

		  (*start).next = temp

		  start = (*start).next

	  }

	  //最后的start 即为最后一个元素 指向头指针
	  (*start) .next = L

	  //返回尾指针
	  return start
  }

  /**
  一般循环链表从tail开始遍历
  循环链表的遍历：与传统的单链表区别为 p.next != nil 改为不等于头指针即可 即又回到头指针了
   */
   func readXhTable(tail *xhTable) {
   		//获取头指针
   		L := (*tail).next
   		h := L //头指针
   		L = h.next  //---第一个元素

   		//从头指针开始 在回来
   		for L.next != h {
   			fmt.Print((*L).data)
   			L = (*L).next
		}

   		//最后一个元素输出
   		fmt.Print((*tail).data)

   }

   /**
   合并2个循环链表  第1个的尾指针指向第2个的第一个元素节点（越过了头指针）第2个的尾指针指向了第一个的头指针
    */

    func mergeXhTable(tail1,tail2 *xhTable) *xhTable{
    	//由于1的尾指针要断裂先记录1的头指针
    	h1 := (*tail1).next

    	//获取2的第一个元素节点
    	node2 := (*tail2).next.next

    	//连接
		(*tail1).next = node2
		(*tail2).next = h1

		//返回尾指针
		return tail2

	}

