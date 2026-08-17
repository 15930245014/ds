# 分组背包
有 N 组物品和一个容量是 V 的背包。每组物品有若干个，**同一组内的物品最多只能选一个**。
每件物品的体积是 cij，价值是 wij，其中 i 是组号，j 是组内编号。
求解将哪些物品装入背包，可使物品总体积不超过背包容量，且总价值最大。输出最大价值。

# 状态转义方程
- dp[i][v] 表示前i组,容量为v时获取的最大价值。
- dp[i][v] = max(dp[i - 1][v],dp[k - 1][v - c[i][j]] + w[i][j] ) j为枚举每组的物品

# 朴素写法 （必会） 
```
dp[0][0] = 0
for i := 0; i < N; i ++ {
    for v := 0; v <=V; v++ { 
        //如果不取当前组
        dp[i][v] = dp[i-1][v]
        
        //取当前组
        for k := 1; k < len(group); k ++ { //循环改组的每一个元素
            //背包容量充足
            if v >= c[i][k] {
                dp[i][v] = max(dp[i][v],dp[i -1][v - c[i][k]] + w[i][k])    
            }
    }
}
return dp[N][V]
```
# 滚动数组写法 -- 推荐
```
dp[0] = 0
for i := 0; i < N; i ++ { //循环每一组
    for v := V; v >= 0;v-- { //循环背包容量 v==0
        //不取当前组,容量为0
        dp[v] = 0
        
        //取当前组
        for k := 1; k < len(group); k ++ { //循环改组的每一个元素
            //背包容量充足
            if v >= c[i][k] {
                dp[v] = max(dp[v],dp[v - c[i][k]] + w[i][k])    
            }
        }
    }
}
return dp[V]
```
# demo
leeCode 1155 https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/
leeCode 2218 https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/description/
2218 其实是一个组里边可以取多个 其实就转化成了多重背包的朴素写法了