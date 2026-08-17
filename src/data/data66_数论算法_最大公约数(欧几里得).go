package main

import "fmt"

/**
	最大公约数： gcd(a,b) 满足 gcd(b,a mod b) （GCD递归定理）
	*欧几里得算法
 */
 func gcdFind(a , b int) int{
 	if b == 0 {
 		return a
	} else {
		return gcdFind(b,a % b)
	}
 }

func main(){
	fmt.Println(gcdFind(16,20))
}