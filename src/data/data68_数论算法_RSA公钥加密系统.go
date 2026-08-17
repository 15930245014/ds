package main

import "fmt"

/**
	RSA (非对称加密算法): 公钥和私钥 加解密可以颠倒顺序----破解方法--整数的因数分解
	1、获取较大的素数 1024位 p,q && p != q
	2、获取n  n = pq
    3、获取fan_N = (q - 1)*(p - 1)
    4、选择e，使得e与fan_N互素，并且小于fan_N, 一般为3，17，
	5、获取d 公式：Φ（n） * k  + 1 =ed      k:为系数，且从1开始，例如：（1,2,3,4,5........）

		公式推出：（Φ（n） * k  + 1 ）/ e = d

		假设k=1时：

		24*k+1/ 5 = d  

		d = 5
		若 k=1 时除不尽的话那就将 k =2 带入计算，还是不行则取 k =3，依次类推。直到（Φ（n） * k  + 1 ）/ e = d 能除尽没有余数为止
	6、RSA公钥 P=(e,n)
	7、RSA私钥 S = (d,n)
*/

//全局变量
var n ,e ,d int

/**
	RSA加密
 */
func RSA(P,Q int)(int,int,int) {
	//获取n
	n = P *  Q

	//获取fan_N
	fanN := (P - 1) * (Q - 1)

	//e一般取3，17等效率最高 e < fanN
	e = 7

	//获取D (fanN*K + 1) / e == d 当能整除时
	k := 1
	for ((fanN * k + 1) % e) != 0 {
		k ++
	}

	//获取d
	d = (fanN * k + 1) / e

	//返回edn
	return e,d,n

}

/**
反复平方法求余
 */
func repeatPowerMode(a,b,mod int) int {
	ans := 1
	for b != 0 {
		if b & 1 == 1 {
			ans = ans * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return ans
}
/**
	RSA加密
 */
func RSACrypt(M int,e,n int) int {//P(e,n) = M^e % n
	return repeatPowerMode(M,e,n)
}

/**
RSA解密
 */
 func RSADecrypt(C,d,n int) int {//S(d,n) =  C^d % n
	 return repeatPowerMode(C,d,n)
 }

func main() {
	//获取edn
	E,D,N := RSA(17,23)

	//明文
	M := 54

	//获取密文
	C := RSACrypt(M,E,N)

	fmt.Println(C)

	//解密
	fmt.Println(RSADecrypt(C,D,N))
}