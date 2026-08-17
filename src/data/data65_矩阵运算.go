package main

import (
	"fmt"
	"math"
)

/**
	矩阵运算(lup分解): Ax = b 其中A为矩阵，一般采用高斯消元法来来求解线性方程组的解
	上三角矩阵（n*n):沿主对角线分割下面的元素全为0
	下三角矩阵(n*n):沿主对角线分割上面的元素全为0

	逆矩阵：A * A(逆) = I (I为方阵且为单位矩阵即主对角线元素权为1) 任何矩阵和单位矩阵相乘均得到矩阵本身

	LUP分解： LUP分解的思想就是找出三个n*n矩阵L、U、P 满足：
				PA = LU  // L为下三角矩阵，U为上三角矩阵 P为置换矩阵
	PAx = Pb => LUx = Pb 令 y = Ux => Ly = Pb  且 A = P(逆)LU

	***本例只讨论A(n*n)情况 ---n行 n个未知数

 */

/**
	矩阵相乘
	2个n阶矩阵相乘
 */
func multiplyMatrix(A [][]float64, B [][]float64) [][]float64 {
	//获取长度
	length := len(A[0])

	//声明返回数组
	var rtMatrix [][]float64

	//初始化返回结果
	for i := 0; i < length; i++ {
		rtMatrix = append(rtMatrix, []float64{})
		for j := 0; j < length; j ++ {
			rtMatrix[i] = append(rtMatrix[i], 0)
		}
	}

	//求解
	for k := 0; k < length; k ++ {
		//A的i行 * B的j列 作为新结果的A对应位置
		for i := 0; i < length; i ++ { //控制B的列
			sum := 0.0
			for j := 0; j < length; j++ { //遍历A的行
				sum = sum + A[i][j]*B[j][k]
			}
			rtMatrix[i][k] = sum
		}
	}

	//返回结果
	return rtMatrix
}

/**
	矩阵加法A,B的阶必须相等
 */
func addMatrix(A, B [][]float64) [][]float64 {
	//矩阵加法对应位置相加即可
	length := len(A[0])

	//声明返回值并初始化
	var rtMatrix [][]float64

	//初始化可以和矩阵相加合并
	for i := 0; i < length; i ++ {
		rtMatrix = append(rtMatrix, []float64{})
		for j := 0; j < length; j ++ {
			rtMatrix[i] = append(rtMatrix[i], 0)
		}
	}

	//矩阵相加
	for i := 0; i < length; i ++ {
		for j := 0; j < length; j ++ {
			rtMatrix[i][j] = A[i][j] + B[i][j]
		}

	}

	//返回
	return rtMatrix
}

/**
		矩阵减法--和加法类似
 */

func decreaseMatrix(A, B [][]float64) [][]float64 {
	//获取长度
	length := len(A[0])

	//声明返回值并初始化
	var rtMatrix [][]float64
	for i := 0; i < length; i ++ {
		rtMatrix = append(rtMatrix, []float64{})
		for j := 0; j < length; j ++ {
			rtMatrix[i] = append(rtMatrix[i], A[i][j]-B[i][j])
		}

	}

	//返回结果
	return rtMatrix
}

/**
	只假设A为方阵
	LU分解 A为方阵
 */
func LUDecompose(A [][]float64)([][]float64,[][]float64) {
	//获取A的长度
	length := len(A)

	//初始化L
	var L [][]float64
	for i := 0; i < length ; i ++ {
		L = append(L,[]float64{})
		for j := 0; j < length;j++ {
			L[i] = append(L[i],0)
		}
		L[i][i] = 1
	}

	//利用高斯消元法
	for i := 0; i < (length - 1); i ++ {//在消元过程中，A共需要消掉n-1个主元下面所有的元素，注意，第n个主元已经是矩阵的最后一个元素了，它的下面和右边都没有其他元素了，所以不存在说对第n个主元下面所有元素消去的情况。这就获得了我们代码的第一个for循环，从第1行主元开始消元，一直到第n-1行主元。
		for j := i + 1; j < length; j ++ { //消去其下的n-1行
			//需要消去元的行所需的乘数
			multiply := A[j][i] / A[i][i]
			L[j][i] = multiply

			//消元
			for k := i ; k < length;k++ {
				A[j][k] -= multiply * A[i][k]
			}

		}

	}

	//返回
	return L,A
}

/**
	矩阵的LUP分解：
	1、解决选主元除数为0的情况
	2、避免除数过小产生的不稳定情况
 */
func LUPDecompose(A [][]float64) ([][]float64,[][]float64,[]int){
	//获取行号
	length := len(A)

	//声明p并初始化---p为置换矩阵
	var p []int
	for i := 0 ; i < length ; i ++ {
		p = append(p,i)
	}

	//初始化L
	var L [][]float64
	for i := 0 ; i < length ; i ++ {
		L = append(L,[]float64{})
		for j := 0; j < length ; j++ {
			L[i] = append(L[i],0)
		}
		L[i][i] = 1 //对角线元素为1
	}


	//LU分解
	for k := 0; k < (length - 1) ; k ++ {
		//声明除数---用于判断除数是否为0
		abs := 0.0
		row := k //默认为下面的第一行
		for m := k; m < length ; m ++ {
			//判断主对角线下行的最大值
			if  math.Abs(A[m][k]) > abs {
				abs = math.Abs(A[m][k])

				//记录行号
				row = m

			}
		}

		//判断
		if abs == 0 {
			fmt.Println("矩阵错误！")
			break
		}

		//如果当前行绝对值不是最大
		if  k != row {
			//交换p[k] 和 p[row]
			p[k],p[row] = p[row],p[k]

		}
		//交换k 和  row 行元素
		for m := 0 ; m < length ; m ++ {
			A[k][m],A[row][m] = A[row][m],A[k][m]
		}

		//LU分解
		for i := k + 1; i < length ; i ++ {//n-1行
			//获取L[i][k] 记录 ---消元法行的乘数
			mul := A[i][k] / A[k][k]
			L[i][k] = mul

			//更新行
			for j := k; j < length; j++ {//A[i][j] - k行*mul
				A[i][j] = A[i][j] - mul * A[k][j]
			}

		}
	}

	//返回
	return L,A,p
}
/**
	pA = LU => Ax = b; pAx = pb;  ->LUx = pb -> Ly = pb -> y = Ux
 */
func LUPSolve(L,U [][]float64,p []int,b[]float64) []float64{
	//获取行
	length := len(L)

	//先求y Ly = pb  L为下三角矩阵 从0遍历
	var x,y []float64
	//初始化下x,y
	for i := 0; i < length; i++ {
		x = append(x,0)
		y = append(y,0)
	}

	for i := 0 ; i < length; i ++ {
		tmp := b[p[i]]
		for j := 0; j < i; j ++ { //这一行
			tmp -= L[i][j] * y[j]
		}
		y[i] = tmp
	}


	//求解 Ux = y u为上三角矩阵
	for i := length - 1; i >= 0; i -- {
		tmp := y[i]
		for j := i + 1; j < length;j ++ {
			tmp -= U[i][j] * x[j]
		}
		x[i] = tmp / U[i][i]

	}

	//打印结果
	return x

}


/**
测试矩阵乘法
 */
func main() {
	var A [][]float64
	A = append(A, []float64{1.0, 2.0, 3.0})
	A = append(A, []float64{4.0, 5.0, 6.0})
	A = append(A, []float64{7.0, 8.0, 0})

	var B [][]float64
	B = append(B, []float64{1, 2, 1})
	B = append(B, []float64{1, 1, 2})
	B = append(B, []float64{2, 1, 1})

	//测试矩阵multiply
	fmt.Println(multiplyMatrix(A, B))
	//[[9 7 8] [21 19 20] [15 22 23]]

	//测试矩阵的加法
	//fmt.Println(addMatrix(A, B))

	//测试矩阵减法
	//fmt.Println(decreaseMatrix(A, B))

	//测试luDecompose
	/*var A1 [][]float64

	A1 = append(A1,[]float64{2,2,2})
	A1 = append(A1,[]float64{4,7,7})
	A1 = append(A1,[]float64{6,18,22})
	L,U := LUDecompose(A1)
	
	//L
	fmt.Println(L)

	//U
	fmt.Println(U)*/

	/*A := [][]float64{}

	A = append(A,[]float64{1,-2,1})
	A = append(A,[]float64{0,2,-8})
	A = append(A,[]float64{5,0,-5})


	B := []float64{0,8,10}

	L,U,P := LUPDecompose(A)

	fmt.Println(LUPSolve(L,U,P,B))*/

}
