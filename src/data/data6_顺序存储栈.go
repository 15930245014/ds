package main

import "fmt"

/**
	栈---后进先出 其实也是一种线性表 其插入（入栈）和删除（出栈） 都是在一端进行的
 */

 /**
 顺序存储栈是利用数组和指向栈顶元素的int 来实现的
  */
  const MAXStackSIZE = 100 //定义常量

  type  stack struct {
  		data [MAXStackSIZE]int
  		top int //top为-1时表示为空栈 ***注意top总是记录栈顶元素下标
  }

  func main() {
  	//构造一个空栈
  	S := createStack()

  	//入栈演示
  	S = pushStack(S,1)
  	S = pushStack(S,2)
  	S = pushStack(S,3)

  	//读取栈
  	readStack(S)

  	//出栈
  	fmt.Print(pullStack(&S))
  	fmt.Print(pullStack(&S))
  	fmt.Print(pullStack(&S))

  }

  /*
  构造一个空栈
   */
  func createStack () stack{
  	//实例化一个栈
  	stack := new(stack)
  	(*stack).top = -1 //此时-1表示为空栈

  	//返回
  	return *stack
  }

  /*
  入栈操作 入栈操作--top++ 新元素放入top位置
   */
  func pushStack(S stack,key int) stack {
  	S.top ++
  	S.data[S.top] = key
  	return S
  }

  /**
  出栈操作每次弹出一个元素---top一开始就 指向栈顶 先出 再top-- 传引用
   */
   func pullStack(S *stack) int {
   		e := S.data[S.top]
   		S.top --
   		return e
   }

   /**
   读取栈中元素
    */
   func readStack(S stack) {
   		for S.top > -1 {
   			fmt.Print(S.data[S.top])
   			S.top --
		}

   }

