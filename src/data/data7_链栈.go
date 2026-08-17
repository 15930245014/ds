package main

import "fmt"

/**
顺序栈无非就是存储形式由数组变成了链表 单链表头指针与栈的top 合二为一就是栈--头插法
 */

 //定义链表
 type lTable struct {
 	data int
 	next *lTable
 }

 //定义链栈---头插法
 type LStack struct {
 	top *lTable //==nil时表示空栈
 	count int //一般定义个数量
 }


func main() {
	//构建链栈
	LS := createLStack()
	//fmt.Print(LS)

	//入栈操作
	LS = pushLStack(LS,1)
	LS = pushLStack(LS,2)
	LS = pushLStack(LS,3)
	printLStack(LS)

	//出栈
	fmt.Print(pullLStack(&LS))
	fmt.Print(pullLStack(&LS))
	fmt.Print(pullLStack(&LS))
}

 /**
 构造
  */
func createLStack() LStack{
	//声明链表
	L := new(lTable)

	//声明链栈
	lStack := new(LStack)

	(*lStack).top = (*L).next

	//返回栈
	return *lStack
}

/**
	入栈操作
 */
 func  pushLStack(lStack LStack,i int) LStack {
 	//构造链表数据
 	 LTable := new(lTable)
	 (*LTable).data = i

	 //判断是否为空栈  top此时指向栈顶 新元素在栈顶
	 LTable.next = lStack.top
	 lStack.top = LTable

	 //返回链栈
	 return lStack
 }

 /**
 	出栈操作
  */
 func pullLStack(S *LStack) int {

 	//获取栈顶元素
 	top := S.top

 	//弹出
	 (*S).top = top.next

	 return  top.data

 }

 /**
 打印链栈
  */
  func printLStack(S LStack) {
  		//获取top
  		top := S.top

  		for top != nil {
  			fmt.Print((*top).data)
  			top = top.next
		}
  }
