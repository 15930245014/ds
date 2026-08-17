package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
	逆波兰公式：（RPN）后缀表达式，不需要括号来实现公式的计算
	1、将公式转化为后缀表达式
	2、后缀表达式求值

	----转后缀表达式思路:
	A、如果是数字直接输出
	B、如果为"("或栈顶元素优先级小于当前入栈
	C、若果为")"则出栈直到碰到第一个"("
	D、若栈顶元素大于当前，则输出直到栈顶元素小于当前或空栈
 */

/**
	定义单链表
 */
type item struct {
	val  byte  //值域
	next *item //指向下一个元素
}

/**
	定义链栈元素结构体
 */
type LineStack struct {
	top    *item //栈顶元素
	length int   //长度
}

/**
	初始化链栈
 */
func initLineStack() LineStack {
	return LineStack{nil, 0}
}

/**
	入栈操作
 */
func pushLineStack(s *LineStack, chr byte) {
	//实例化结点
	node := new(item)
	node.val = chr

	//写入栈顶
	node.next = s.top
	s.top = node

	//栈长度+1
	s.length ++
}

/**
出栈操作
 */
func popLineStack(s *LineStack) byte {

	//长度-1
	s.length --
	rt := s.top.val

	//栈顶元素下移
	s.top = s.top.next

	//返回
	return rt

}

/**
	获取后缀表达式
	@params s string 公式
	@return string 后缀表达式
	@author 许校
 */
func getSuffixFormula(s string) []string {
	//去除空格 "a + b" => "a+b"
	s = strings.Replace(s, " ", "", -1)

	//定义数字字典
	numMap := make(map[byte]int)

	numMap['0'] = 1
	numMap['1'] = 1
	numMap['2'] = 1
	numMap['3'] = 1
	numMap['4'] = 1
	numMap['5'] = 1
	numMap['6'] = 1
	numMap['7'] = 1
	numMap['8'] = 1
	numMap['9'] = 1

	//定义符号优先级
	operaMap := make(map[byte]int)

	operaMap['+'] = 1
	operaMap['-'] = 1
	operaMap['*'] = 2
	operaMap['/'] = 2
	operaMap['('] = 3
	operaMap[')'] = 3

	//声明返回切片
	var rtArray []string

	//声明一个空栈
	stack := initLineStack()

	//遍历字符串
	for i := 0; i < len(s); i++ {
		//获取字符
		v := s[i]

		//判断是否为数字
		if numMap[v] == 1 {
			tmp := string(v)
			for j := i + 1 ; j < len(s) && numMap[s[j]] == 1; j ++ {
				tmp += string(s[j])
				i++
			}

			rtArray = append(rtArray, tmp)
			continue
		}

		//判断是否为左括号
		if v == '(' || stack.length <= 0 {
			//入栈
			pushLineStack(&stack, v)
			continue
		}

		//判断当前是否为右括号
		if v == ')' {
			//出栈直到遇到"("
			for stack.length > 0 && stack.top.val != '(' {
				//出栈
				rtArray = append(rtArray,string(popLineStack(&stack)))
			}
			//去除最后的"("
			popLineStack(&stack)
			continue

		}
		//判断当前和栈顶元素的优先级
		if operaMap[v] > operaMap[stack.top.val] {
			pushLineStack(&stack, v)

		} else { //当前元素优先级小于栈顶
			for stack.length > 0 && operaMap[stack.top.val] >= operaMap[v] && stack.top.val!='(' {
				rtArray = append(rtArray,string(popLineStack(&stack)))
			}

			//当前入栈
			pushLineStack(&stack, v)
		}

	}

	//判断栈是否为空
	if stack.length > 0 { //依次弹出
		for stack.length > 0 {
			rtArray = append(rtArray,string(popLineStack(&stack)))
		}
	}

	//返回
	return rtArray

}
/**
	实现逆波兰公式
 */
func myRpn(s string) int{
	//获取后缀表达式
	suffixArray := getSuffixFormula(s)
	fmt.Println(suffixArray)

	//定义操作符以及参数个数
	operator := make(map[string]int)
	operator["+"] = 2
	operator["-"] = 2
	operator["*"] = 2
	operator["/"] = 2

	//定义操作符函数
	operaFunc := make(map[string]func(a,b int)int)

	//采用匿名函数
	operaFunc["+"] = func(a,b int) int {
		return a + b
	}
	operaFunc["-"] = func(a,b int) int {
		return a - b
	}
	operaFunc["*"] = func(a,b int) int {
		return a * b
	}
	operaFunc["/"] = func(a,b int) int {
		return a / b
	}

	//定义栈--存放结果
	var stack []int

	//遍历后缀表达式
	for i := 0; i < len(suffixArray); i++ {
		//获取元素
		v := suffixArray[i]

		//判断是否是数字
		if operator[v] > 0 {
			//获取参数个数---利用反射去执行 ---本例默认为2个参数
			length := len(stack)
			num1 := stack[length - 1]
			num2 := stack[length - 2]

			//出栈2个
			stack = stack[0:length - 2]

			//计算并入栈
			stack = append(stack,operaFunc[v](num2,num1))

		} else {//表示为数字--入栈
			num,_ := strconv.Atoi(v)
			stack = append(stack,num)
		}

	}

	//返回值
	return stack[0]
}
func main() {
	//测试生成后缀表达式
	s := "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(myRpn(s))
}
