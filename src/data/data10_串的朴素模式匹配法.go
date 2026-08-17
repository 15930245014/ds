package main

import "fmt"

/**
	串--就是所谓的字符串串中的每个元素都是字符 ==正因为每个元素为字符故链式存储方式占用内存太大 且不如顺序存储灵活
	串的存储方式有顺序存储方式和链式存储-----主要采用顺序存储的方式且字符串的结束一般以\0结束
	2个串的比较是从两个数组从0~\0之间顺序比较ascii 只要分出大小终止
 */

 /**
 	串的模式匹配 值得是主串是否包含子串 并返回其所在的小标 比如 "hello" 包含h 返回0， -1表示不存在
  */

  func main() {
  	//测试我的方法----厉害厉害
  	fmt.Print(myIndex("helloword","word"))
  	fmt.Print(myIndex("helloword","word!"))

  	//测试朴素模式匹配法---思路和我的异曲同工代码略简单
	  fmt.Print(simpleIndex("helloword","word"))
	  fmt.Print(simpleIndex("helloword","word!"))
  }
  /**
  返回s2在s1出现的位置 ---我之前的写法
   */
  func myIndex(s1,s2 string) int {
  	//定义2个其实下表
  	i := 0

  	rtIndex := -1

  	for i < len(s1) {
  		//判断是否等于s2头元素
  		if s1[i] == s2[0] {
  			j := 1
  			for j < len(s2) && i + j < len(s1) && s2[j] == s1[i + j] { //i+j < len(s1) s1有越界的情况
  				j ++

			}

  			//判断j是否等于len(s2)
  			if j == len(s2) {
  				rtIndex = i
  				break
			}

		}
  		//继续遍历
  		i ++


	}

  	//返回索引
  	return rtIndex

  }
  /**
	朴素模式匹配发
  */
  func simpleIndex(s1,s2 string) int {
  	//定义下标
  	i,j,rtIndex := 0,0,-1

  	//遍历条件为都未达到瓶颈
  	for i < len(s1) && j < len(s2) {
  		//判断是否相等
  		if s1[i] == s2[j] {
  			i++
  			j++

		} else {
  			i = i - j + 1 //回到原来的位置 在后移1为
  			j = 0 //j回到原点

		}

	}

  	//判断j 最后是否走完 走完说明完全匹配
  	if j == len(s2) {
  		rtIndex = i - j //最开始位置
	}

  	//返回结果
  	return rtIndex

  }
