package main

import "fmt"

//问题描述
/**
给定一个字符串，你请找出其中不含有重复字符的  最长子串 的长度。
输入： “abcabcbb”
 输出： 3
 解释：因为无重复字符的最长子串是"abc"，所以其长度为3。
 */

 func main() {
	fmt.Print(getMaxChildString("abcabcbb"))
 }
/**
优化方案
 */
 func getMaxChildString(s string) int {
 	//构造起止 下标 end
 	end := 0

 	//记录byte
 	var btArray []byte

 	max := 0

 	for end < len(s) {

 		if len(btArray) == 0 || findIndex(s[end],btArray) == -1 {
 			//顺序存入切片
 			btArray = append(btArray,s[end])
 			end ++

		} else {
 			if len(btArray) > max {
 				max = len(btArray)
			}
 			//获取索引
 			index := findIndex(s[end],btArray)
 			btArray = btArray[index+1:]

		}

	}

 	return max

 }
 /**
 辅助函数
  */

/**
查询在元素在切片的位置
*/
func findIndex(key byte,s []byte) int{

	for i,value := range s {
		if value == key {
			return i
			break
		}
	}

	return -1
}