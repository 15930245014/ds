package main

import (
	"fmt"
)

/**
PMP模式匹配法：abcadss 中存在 bcd  由于bca a与bcd最后一个相同 故i 不回溯 只是让a与bcd的最后一个比较
详见数据结构讲解

 */
func main() {
	fmt.Print(KMPIndex("helloword","word"))
}
/**
	构造next数组

 */
func makeNextArray(s string)[20]int {
	//声明next数组
	var next [20]int

	//next数组0为0
	next[0] = -1


	//定义开始和结束
	begin,end := -1,0

	for end < len(s) - 1  {

		if begin == -1 || s[begin] == s[end] {
			begin++
			end++
			//next[end] = begin 原始的
			/***改进**/
			if s[begin] != s[end] {
				next[end] = begin
			} else {
				next[end] = next[begin]
			}

		} else {
			begin = next[begin]
		}

	}

	return next

}

/**
   朴素模式匹配发
 */
func KMPIndex(s1,s2 string) int {
	//定义下标
	i,j,rtIndex := 0,0,-1

	//获取next
	next := makeNextArray(s2)

	//遍历条件为都未达到瓶颈
	for i < len(s1) && j < len(s2) {
		//判断是否相等
		if j == -1 || s1[i] == s2[j] {
			i++
			j++

		} else {

			j = next[j]

		}

	}

	//判断j 最后是否走完 走完说明完全匹配
	if j == len(s2) {
		rtIndex = i - j //最开始位置
	}

	//返回结果
	return rtIndex

}

