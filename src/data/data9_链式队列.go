package main

import "fmt"

/**
	链式队列，---基于单链表改进   ****其实队列的整体方向是向右移动的
 */

 /**
 	声明单链表
  */
  type lineTable struct {
  		data int
  		next *lineTable
  }

  /**
   声明链式队列
   */
   type lineQueue struct {
   		top *lineTable   //指向头指针 ---栈的top指向第一个元素
   		tail * lineTable  //指向最后一个元素
   }

   func main() {
   		//测试队列
   		lineQueue := makeLineQueue()

   		//入队
   		lineQueue = pushLineQueue(lineQueue,1)
   		lineQueue = pushLineQueue(lineQueue,2)
   		lineQueue = pushLineQueue(lineQueue,3)
   		lineQueue = pushLineQueue(lineQueue,4)

   		//出队操作
   		fmt.Println(pullLineQueue(&lineQueue))
   		fmt.Println(pullLineQueue(&lineQueue))
   		fmt.Println(pullLineQueue(&lineQueue))
   		fmt.Println(pullLineQueue(&lineQueue))
   		fmt.Println(pullLineQueue(&lineQueue))

   		//入队
	   lineQueue = pushLineQueue(lineQueue,5)

	   fmt.Println(pullLineQueue(&lineQueue))
	   fmt.Println(pullLineQueue(&lineQueue))

   }

   /**
   		构造空队列tail& top = 头指针
    */
   func makeLineQueue() lineQueue{
   		//声明头指针
   		//L := new(lineTable)

   		//构造空队列
   		line := lineQueue{nil,nil}

   		//返回空队列
   		return line

   }

   /**
   	 入队操作 ---尾指针入队====不存在满的情况
    */
   func pushLineQueue(line lineQueue,key int) lineQueue {

   		//构造新的结点
   		newNode := new(lineTable)

	   (*newNode).data = key
	   (*newNode).next = nil

	   //判断是否为空
	   if line.tail == nil {
	   	    line.tail = newNode
	   		line.top = newNode
	   } else {
		   //插入元素
		   (*(line.tail)).next = newNode
	   }

   		//尾指针指向最后一个元素
   		line.tail = newNode

   		//返回
   		return line
   }
	/**
		 出队操作 ---头指针出队 top 永远指向头指针
	 */
	func pullLineQueue(line *lineQueue) int{

		//判断是否为空
		if line.top == nil {
			fmt.Println("队列为空!")
			return -1;
		}

		//获取top
		top := line.top

		//更新头
		line.top = top.next

		//判断是否为空
		if line.top == nil {
			//更新尾
			line.tail = nil

		}

		//返回
		return top.data
	}