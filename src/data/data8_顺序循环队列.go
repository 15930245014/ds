package main

import "fmt"

/**
队列（Queue）:队列(Queue)是只允许在一端进行插入操作，而在另一端进行删除操作的线性表。
允许插入的端是队尾（tail）入队，
允许删除的端是队头(top)。 出队
特点：先进先出
 */

 /**
 	传统的实现
 	顺序存储就是用数组实现，比如有一个ｎ个元素的队列，数组下标0的一端是队头，入队操作就是通过数组下标一个个顺序追加，不需要移动元素
   但是如果删除队头元素，后面的元素就要往前移动，对应的时间复杂度就是O(n)，性能自然不高。
  */

  /**
  为了提高出队的性能，就有了循环队列，
  什么是循环队列呢？就是有两个指针，
  top指向队头，tail指向对尾元素的下一个位置，
  元素出队时top往后移动，如果到了对尾则转到头部，同理入队时tail后移，如果到了对尾则转到头部，
  这样通过下标top出队时，就不需要移动元素了。
   */

   /**
   		top和tail 从0开始
   		循环队列：为空时 即taii == top
   		队列满时：tail + 1 == top(绕一圈回来)但考虑top为0 tail 为maxSize 情况 则有 (tail - top) ==maxSize 合并为（tail + 1） % maxSize == top
   		队中元素个数：tail - top || maxSize - top + tail   合并  (tail - top + maxSize) % maxSize
    */
    const MaxQuenu  = 5

    //循环链表的结构体
    type Queue struct {
    	data [MaxQuenu]int //存放数值的数组
    	top int //类似头指针
    	tail int //类似尾指针
	}


func main() {
	//构造空队列
	queue := makeQueueEmpty()

	//入队测试 ----初始化是 tail 值为0
	queue = pushQueue(queue,1)
	queue = pushQueue(queue,2)
	queue = pushQueue(queue,3)
	queue = pushQueue(queue,4)

	readQueue(queue)  //0，1，2，3，4

	//出队操作
	fmt.Println("")
	fmt.Println(pullQueue(&queue))
	fmt.Println(pullQueue(&queue))
	fmt.Println(pullQueue(&queue))
	fmt.Println(pullQueue(&queue))
	fmt.Println(pullQueue(&queue))

	queue = pushQueue(queue,1)
	queue = pushQueue(queue,2)
	queue = pushQueue(queue,3)
	queue = pushQueue(queue,4)
	readQueue(queue)

	fmt.Println(pullQueue(&queue))
	fmt.Println("")
	queue = pushQueue(queue,5)
	readQueue(queue)






}
/**
	构造空队列
 */
 func makeQueueEmpty() Queue{
 	//top和tail 相同并且同时为0
 	return Queue{[MaxQuenu]int{},0,0}
 }

 /**
 	入队操作==从尾指针tail开始入队
  */
  func pushQueue(queue Queue,key int) Queue{
  	//获取头和尾
  	top := queue.top
  	tail := queue.tail

  	//判断是否满了
  	if (tail + 1) % MaxQuenu == top {
  		return queue
	}


  	//尾指针后移如果越界则从0开始
  	newTail := (tail + 1) % MaxQuenu

  	//存放数据
  	queue.data[newTail] = key

  	//修改tail
  	queue.tail = newTail

  	//返回
  	return queue
  }

  /**
  	出队操作 ==从头开始
   */
   func pullQueue(queue *Queue) int {
   		//获取top和tail
   		top := (*queue).top
   		tail := (*queue).tail

   		//判断是否为空
   		if top > tail { //刚开始相等时
   			return 100000000000  //应当抛错
		}

   		//获取元素
   		e := (*queue).data[top]

   		//top 右移
   		newTop := (top + 1) % MaxQuenu

   		//修改top
	   (*queue).top = newTop

	   //返回
	   return e

   }

   //读取循环队列
   func readQueue(queue Queue) {
   	//获取tail和top
   	top := queue.top
   	tail := queue.tail

   	for top % MaxQuenu != tail {
   		fmt.Print(queue.data[top % MaxQuenu])
   		top ++
	}

   	//弹出同时指向的元素
   	 fmt.Print(queue.data[top % MaxQuenu])

   }
