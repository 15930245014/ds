package 算法

/*
可以使用经典的 埃拉托斯特尼筛法（Sieve of Eratosthenes）埃氏素数筛。
这个算法的时间复杂度是 𝑂(𝑛log⁡log⁡𝑛)，非常高效。它的核心思想是：
从最小的质数 2 开始，把所有 2 的倍数标记为非质数，然后继续标记 3 的倍数，依此类推 直到 i *i <=n。
*/
func SieveOfEratosthenes(n int) []int {
	// 创建一个布尔切片，表示每个数是否为质数
	// 默认认为每个数都是质数（true），0和1除外
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// 埃氏素数筛
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// 将 i 的倍数标记为非质数
			for j := i * i; j <= n; j += i { //记住这个
				isPrime[j] = false
			}
		}
	}

	// 收集所有质数
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
