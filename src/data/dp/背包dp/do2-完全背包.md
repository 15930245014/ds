    # 完全背包
有N种物品和一个容量为V的背包，每种物品都有无限件可用。第i种物品的费用是c[i]，价值是w[i]。求解将哪些物品装入背包可使这些物品的费用总和不超过背包容量，且价值总和最大。

- dp[i][v] 表示前i件物品放入容量为j的最大价值
- 状态转义方程: dp[i][v] = max(dp[i - 1][v - k * c[i]] + k * w[i])(0<=k*c[i]<= v)
- k取0表示不拿 

# 优化完伪代码
最大值或最小值
```
dp[0...N] = 0
for i =1...N {
    for v=c[i]..V {
        dp[v] = max(dp[v],dp[v - c[i]] + w[i]) //都可以放得下
    }
}
return dp[N]
```
最大值最小值在处理“恰好”问题上 
```
dp[0] = 0 
dp[1...N] = +∞或-∞
for i =1...N {
    for v=c[i]..V {
        if dp[v- c[i]] ==  +∞或-∞ {continue}  //不能恰好放的下
        dp[v] = max(dp[v],dp[v - c[i]] + w[i])
    }
}
if dp[v] == +∞或-∞ {return ...} //不满足条件的返回值判断
return dp[N]
```
在方案总数上
```
    dp[0] = 1
    dp[1..N] = 0
    for i=1...N {
        for v = c[i]...V{
            dp[v] += dp[v - c[i]]
        }
    }
    return dp[V]
```
# 题干相关
求组合数，遍历顺序：先遍历物品再遍历背包容量
https://leetcode.cn/submissions/detail/578930138/
```
    //求组合数
func change(amount int, coins []int) int {
    //完全背包 恰好问题
	dp := make([]int,amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
```
求排列数，遍历顺序：先遍历背包容量再遍历物品
https://leetcode.cn/problems/combination-sum-iv/description/
```
//求排列数
func combinationSum4(nums []int, target int) int {
	//完全背包恰好问题
	dp := make(map[int]int)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j:= 0; j < len(nums);j ++ {
			//判断
			if i >= nums[j]  {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
```
