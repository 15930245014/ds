package main

import (
    "fmt"
)

/**
	描述:
给定一个字符串s，找到s中最长的回文子串。你可以假设  s的最大长度为1000。

示例1：

输入： “babad”
 输出： “bab”
 注意： “aba”也是一个有效答案。
示例2：

输入： “acbbcd”
 输出： “bb”
 */

 func main () {
 	fmt.Println(longestPalindrome2("acbcd"))
 }
 /**
 	dp[i][j] = dp[i + 1][j - 1] if s[i] == s[j]  ==3字母以上
    dp[i][j] = s[i]   s[i] == s[j] 1字母
    dp[i][j] = dp[i][i]+s[j]   2字母

  */
 func longestPalindrome(s string) string {
    //获取字符串长度
    length := len(s) - 1

     //判断字符串
    if len(s) == 0 {
        return s
    }

    //定义返回
    rtStr := string(s[0])

    //初始化dp
    dp := map[string]bool{}
    for i :=0; i <= length; i ++ {
        dp[string(s[i])] = true

    }

    //2字母初始化
    for i := 0; i <= length - 1; i ++ {
        if s[i] == s[i + 1] {
            dp[string(s[i:i+2])] = true
            rtStr = string(s[i : i+2])
        }
    }

    //遍历
    for i := length - 1; i >= 0; i-- {
        for j := i + 1; j <= length;j++ {
            //判断是否相等
            if s[i] == s[j] && dp[string(s[i + 1:j])] == true {
                //判断
                if len(rtStr) < (j - i + 1) {
                    rtStr = string(s[i:j + 1])
                }

                dp[string(s[i:j+1])] = true
            }
        }

    }
    //返回
    return rtStr
 }
 /**
    回文子串
  */
func longestPalindrome2(s string) string {
    //获取字符串长度
    length := len(s) - 1

    //判断字符串
    if len(s) == 0 {
        return s
    }

    //定义返回
    rtStr := string(s[0])

    //初始化dp
    dp := map[string]bool{}
    for i :=0; i <= length; i ++ {
        dp[string(s[i])] = true

    }
    dp[""] = true

    for i := 1; i <= length; i++ {
        for k := 0; k < i; k++ {
            //判断是否相等
            if s[k] == s[i] && dp[string(s[k+1:i])] == true  {
                if len(rtStr) < ( i - k + 1) {
                    rtStr = string(s[k:i + 1])
                }

                dp[string(s[k :i + 1])] = true


            } else {
                //写入缓存
                dp[string(s[k:i + 1])] = false
            }

        }
    }

    //返回
    return rtStr

}
