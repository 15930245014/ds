package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	整数的因式分解：就是分解为2个素数的积。
	pollard 的rho启发式法：（波拉德的rho启发式法） ----膜拜吧
 */

/**
	辅助函数：求最大公约数--欧几里得算法
 */
func myGcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return myGcd(b, a % b)
	}
}

/**
	波拉德的rho启发式因子分解
 */
func pollardRho(n int) (int,int) {
	//索引
	i := 0

	//定义x切片
	var x []int

	//随机数x[i]
	x = append(x, myRand(0, n-1))
	y := x[i]
	k := 2
	var d int

	for true {
		i ++
		x = append(x,(x[i - 1] * x[i - 1] - 1) % n)
		d = myGcd(y - x[i],n)

		if d != -1 && d != 1 && d != n {
			break
		}
		if i == k {
			y = x[i]
			k = k * 2
		}

	}

	//可能会小于0
	if d < 0 {
		d *= -1
	}

	//返回
	return d,n / d

}

/**
	随机数
 */
func myRand(min, max int) int {
	//设置
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func main() {
	fmt.Println(pollardRho(31 * 23))
}