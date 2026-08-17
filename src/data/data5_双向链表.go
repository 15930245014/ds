package main

import "fmt"

/**
双向链表 ：由于单向链表不可逆 即只能从头指针开始遍历 双向链表多了一个前驱指针来记录上一个节点位置
 */

 /**
  双向链表的数据结构
  */
  type twoWayTable struct {
  	 data int
  	 pre *twoWayTable //前驱
  	 next *twoWayTable //后继
  }

  func main() {
  	//构造元素
  	arr := []int{1,2,3}

  	//构造头指针
  	L := new(twoWayTable)

  	//头插法构造双向链表
  	tail :=  createTwoWayTableHeader(L,arr)

  	//正序读取
  	readTwoWayTable(L)

  	//逆序读取
  	readTwoWayTableDesc(tail)

  	fmt.Print("\n")

  	//尾插法构造双向链表
  	L2 := new(twoWayTable)
  	tail2 := createTwoWayTableTail(L2,arr)

  	//正序读取
  	readTwoWayTable(L2)

  	//逆序读取
  	readTwoWayTableDesc(tail2)

  	fmt.Println("")

  	//添加元素
  	addTwoWayTable(L2,2,0)
  	readTwoWayTable(L2)
  	fmt.Println("")

  	//删除元素
  	deleteTwoWayTable(L2,2)
  	readTwoWayTable(L2)


  }

  /**
  	正序遍历链表元素---从第一个值域节点开始
   */
  func readTwoWayTable(L *twoWayTable) {
  	p := (*L).next

  	for p != nil {
		fmt.Print((*p).data)
		p = (*p).next

	}

  }
  /**
  	倒序遍历链表元素
   */
  func readTwoWayTableDesc(tail *twoWayTable) {
  	p:= tail
  	for p != nil {
  		fmt.Print(p.data)
  		p = p.pre
	}
  }

  /**
   头插法---构造双向链表
   */
   func createTwoWayTableHeader(L *twoWayTable,arr []int) *twoWayTable{
   		//头插法
   		var tail *twoWayTable
   		for _,value := range arr {
   			//声明临时双向链表
   			twoWay := new(twoWayTable)
			(*twoWay).data = value

			Lnext := (*L).next

			(*twoWay).next = (*L).next

			(*L).next = twoWay

			if Lnext != nil {
				(*Lnext).pre = twoWay

			} else {//最后一个节点
				tail = twoWay
			}
		}

   		return tail
   }

   /*
   尾插法---构造双向链表
    */
    func createTwoWayTableTail(L *twoWayTable,arr []int) *twoWayTable {
    	//构造第一个节点start
    	start := new (twoWayTable)
    	(*start).data = arr[0]
		(*start).pre = L
		(*L).next = start

		var tail *twoWayTable

		//插入其他的双向链表
		for j := 1;j < len(arr);j++ {
			temp := new(twoWayTable)

			//放入值域
			(*temp).data = arr[j]

			(*temp).pre = start
			(*start).next = temp

			start = (*start).next

			if j == len(arr) - 1 {
				(*temp).next = nil
				tail = temp
			}

		}

		//返回最后一个元素
		return  tail

	}

    //双向链表的插入操作
    func addTwoWayTable(L *twoWayTable,i,key int) {
		//丝路先寻找i-1位置
		p := (*L).next
		j := 1

		//寻找i-1节点
		for p != nil &&  j < i - 1 {
			p = p.next
			j++
		}

		node := p

		//构造要插入的节点
		newNode := new (twoWayTable)
		(*newNode).data = key

		//执行插入操作
		nextNode := (*node).next

		if nextNode != nil {
			(*nextNode).pre = newNode
		}
		(*node).next = newNode
		(*newNode).pre = node
		newNode.next = nextNode
	}

//双向链表的删除操作
func deleteTwoWayTable(L *twoWayTable,i int) {
	//丝路先寻找i-1位置
	p := (*L).next
	j := 1

	//寻找i-1节点
	for p != nil &&  j < i - 1 {
		p = p.next
		j++
	}

	node := p

	deleteNode := (*node).next
	(*node).next = (*deleteNode).next
	(*(*deleteNode).next).pre = node
}

