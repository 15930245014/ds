package main

import (
	"fmt"
	"math"
)

/***
	线性规划：线性规划问题又称线性规划，在数学中线性规划（Linear Programming，简称LP）特指目标函数和约束条件皆为线性的最优化问题。
	标准型：描述线性规划问题的常用和最直观形式是标准型如：
	max 2 *x1 + 3 * x2
	x1 + x2 <= 10
	2 * x1 +  3 *x2 <= 12
	x1,x2 >= 0
　　如上即为标准型例子

	松弛型：通过添加增广参数来使不等式成立，上述标准型转松弛型
	Z = max(2 * x1 + 3 * x2)
	x1 + x2 + x3 = 10
   2x1 + 3x2 + x4 = 12
	x1,x2,x3,x4 >= 0
		*x3,x4又叫做松弛变量

	单纯形法解线性规划：一种古老但很有效的解决线性规划方法，类似的可以用图解法解
	涉及概念：
	目标函数(C)：即要求解的目标值函数公式
	基变量：等式左边的变量叫做基变量
	非基变量：等式右边的叫做非基变量
	检验数：非基变量的目标函数值 - （基变量目标值* 非基变量系树） 非基变量检验数存在 >0 时 一般还未达到最优解 决策时检验数最大的非基变量要通过转动来变为基变量
	基变量转出原则：选取bi / Aj 最小的作为转出变量  当检验数>0 Aj全小于0时 没有可行解
 */
 func initSimpleLp(A [][]float64,b,c []float64)([]int,[]int,[][]float64,[]float64,[]float64,[]float64)  {
	//初步判断是否有可行解-- 增广矩阵列最小值小于0
	if minCheckNum(b) < 0 {
		fmt.Println("无可行解!")
	}
	//A(m * n)
	m := len(A)     //矩阵行数
	n := len(A[0])  //未知数个数

	//将标准型化作松弛型---每行添加一个松弛变量下标为n+i (1<=i<=m)
	for i := 1; i <= m; i ++ {
		//更行目标函数 --目标函数的松弛变量下标均为0
		c = append(c,0)

		//添加如矩阵A
		for j := 0; j < m; j ++ {//未知数个数扩展到n+m个
			A[i - 1] = append(A[i - 1],0)
		}
		A[i - 1][n + i - 1] = 1
	}

	//声明基变量和非基变量矩阵
	var B,N []int

	//初始化基变量
	for i := 1; i <= m ; i++ {
		B = append(B,n + i - 1)
	}

	//初始化非基变量
	for i := 0; i < n ; i ++ {
		N = append(N,i)
	}

	//初始化检验数
	var checkArr []float64
	for i := 0 ; i < (m + n); i ++ {
		checkArr = append(checkArr,0)
	}

	//求解检验数--根据非基变量
	for i := 0; i < len(N); i ++ {
		//获得非基变量下标
		n := N[i]

		//获取检验数
		num := c[n]
		for j := 0; j < len(B); j ++ {
			num = num - c[B[j]] * A[j][n]
		}

		//写入
		checkArr[n] = num
	}

	//返回初始值
	return N,B,A,b,c,checkArr
 }
 /**
 	单纯形算法
  */
 func simpleLp(A [][]float64,b,c []float64) float64 {
 	//初始化
	 N,B,A,b,c,checkArr := initSimpleLp(A,b,c)

	 //最大检验数
	 maxCheck := maxCheckNum(checkArr)

	 //当检验数最大值存在
	 for maxCheck > 0 {
	 	//获取非基变量转入基变量的下标
	 	n := 0
	 	for  n < len(checkArr) {
	 		if maxCheck == checkArr[n] {
	 			break
			}
	 		n++
		}

	 	//判断需要转出的基变量min(bj / Ajn)
	 	base := -1
	 	min := math.MaxFloat64

	 	for j := 0; j < len(B); j ++ {
	 		//判断
	 		if A[j][n] <= 0 {
	 			continue
			}

	 		//获取
	 		tmp := b[j] / A[j][n]

	 		if tmp < min {
	 			min = tmp
				base = j
			}
		}

	 	//判断
	 	if base == -1 && min == math.MaxFloat64 {
	 		fmt.Println("无可行解!")
	 		break
		}

	 	//找到转入转出变量n,base--旋转操作
		 N,B,A,b,c,checkArr = simplePivot(N,B,A,b,c,base,n)

		 //更新检验数
		 maxCheck = maxCheckNum(checkArr)

	 }

	 //返回最优值
	 Z := 0.0
	 for i := 0; i < len(B); i ++ {
	 	Z = Z + c[B[i]] * b[i]
	 }

	 //返回
	 return Z


 }
 /**
 	*@param N 非基变量
 	*@param B 基变量
 	*@param A 增广矩阵
 	*@param b 约束条件
 	*@param c 目标函数
 	*@param l 转出变量
 	*param e 转入变量
  */
 func simplePivot(N,B []int,A [][]float64,b []float64,c []float64,l,e int)([]int,[]int,[][]float64,[]float64,[]float64,[]float64 ){
 	//更新转出变量那一行 A[l][o] = 1 行系树同时除以 A[l][o]
 	k := A[l][e]
 	for j := 0; j < len(A[l]); j ++ {
 		A[l][j] /= k
	}
 	//更新b
 	b[l] = b[l] / k

 	//其他行利用高斯消元法转为0
 	for i := 0; i < len(B); i ++ {
 		if i == l {
 			continue
		}

 		//初等行变换
 		bk := A[i][e] / 1
 		b[i] = b[i] - bk * b[l]
 		for j := 0; j < len(A[i]); j ++ {
 			A[i][j] = A[i][j] -  bk * A[l][j]
		}
	}

 	//更新基变量和非基变量
 	B[l],N[e] = N[e],B[l]

 	//获取检验数
	 var checkArr []float64
	 for i := 0 ; i < (len(B) + len(N)); i ++ {
		 checkArr = append(checkArr,0)
	 }

	 //求解检验数--根据非基变量
	 for i := 0; i < len(N); i ++ {
		 //获得非基变量下标
		 n := N[i]

		 //获取检验数
		 num := c[n]
		 for j := 0; j < len(B); j ++ {
			 num = num - c[B[j]] * A[j][n]
		 }

		 //写入
		 checkArr[n] = num
	 }

	 //返回
	 return N,B,A,b,c,checkArr

 }

 /**
 	获取最小值
  */
 func minCheckNum(numArr []float64) float64 {
 	//初始
 	min := numArr[0]

 	//比较
 	for i := 1; i < len(numArr); i ++ {
 		if min > numArr[i] {
 			min = numArr[i]
		}
	}

 	//返回
 	return min

 }
 /**
 	获取最大值
  */
 func maxCheckNum(numArr []float64) float64 {
 	//初始
 	max := numArr[0]

 	//比较
 	for i := 1; i < len(numArr);i ++ {
 		if max < numArr[i] {
 			max = numArr[i]
		}
	}

 	//返回
 	return max
 }

func main() {
/**
	max z = 3 * x1 + 2 * x2
			2 * x1 + x2 <= 10
			x1 + x2 <= 8
	x1,x2 >= 0
 */
 var A [][]float64
 A = append(A,[]float64{2,1})
 A = append(A,[]float64{1,1})
 //A = append(A,[]float64{3,2})

 b := []float64{10,8}
 c := []float64{3,2}

 fmt.Println(simpleLp(A,b,c)) //18 x1=2 x2=6

}

