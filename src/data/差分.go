package diffArr

import "fmt"

func diffArr() {
	/**
	创建差分数组，假设原来是A,差分数组为D 则有：
	D[0] = A[0]
	D[i] = A[i] - A[i-1] (i= 1 to N)
	*/
	A := []int{1, 2, 3, 4, 5}
	n := len(A)
	D := make([]int, n)
	D[0] = A[0]
	for i := 1; i < n; i++ {
		D[i] = A[i] - A[i-1]
	}

	/**
	区间增加：[L,R] 增加x
	D[L] += x
	if R+1 < n {
		D[R] -= x
	}

	区间减少x
	D[L] -=x
	if R+1 < n {
		D[R] += x
	}
	*/
	// 对区间 [1, 3] 减去 5
	L1, R1, x1 := 1, 3, 5
	D[L1-1] -= x1
	if R1 < n {
		D[R1] += x1
	}

	//回复数组
	R := make([]int, n)
	R[0] = D[0]
	for i := 1; i < n; i++ {
		R[i] = R[i-1] + D[i]
	}
	fmt.Println(R)
}
