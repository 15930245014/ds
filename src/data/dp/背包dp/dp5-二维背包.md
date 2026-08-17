# 二维背包问题
- 二维费用的背包问题是指对于每件物品，具有两种不同的费用，选择这件物品必须同时付出这两种代价，对于每种代价都有一个可付出的最大值（背包容量），求选择物品可以得到最大的价值。
- 例如，有一个背包，它的容量为V，它的重量限制为U。有N件物品，第i件物品的体积为a[i]，重量为b[i]，价值为c[i]。在不超过背包容量和重量限制的情况下，如何选择物品使得背包内物品的总价值最大？

# 状态转义方程
- 费用加了一维，只需状态也加一维即可
- dp[i][v][u] 表示前i件物品放入容量为v,重量为u的获取的最大值
- 状态转义方程：dp[i][v][u] = max(dp[i - 1][v][u],dp[i - 1][v-a[i]][u - b[i]] + w[i])

# 朴素实现
```
    for i := 0; i < N; i ++ {
        for v := 0;v <= V; v ++ {
            for u:=0;u <=U;u++ {
                if v < a[i] || u <b[i] { //放不下
                      dp[i][v][u] = dp[i - 1][v][u]
                      continue
                }
                //取或不取
                dp[i][v][u] = max(dp[i][v][u],dp[i - 1][v - a[i][u - b[i]] + w[i])
            }       
        }
    }
    return dp[N][V][U]
```

# 滚动数组实现0-1背包
```
dp[0][0] = 0
for i := 0; i < N; i ++ {
    for v := V; v >= a[i];v-- {
        for u:=U;u >= u[i];u-- {
            dp[v][u] = max(dp[v][u],dp[v-a[i]][u-b[i]] + w[i])
        }
    }
}
return dp[V][U]
```

# 滚动数组 完全背包
```
    dp[0][0] = 0
    for i :=0; i < N; i ++ {
        for v := a[i];v <=V;v++ {
            for u:= b[i]; u <=U;u++ {
                dp[v][u] = max(dp[v][u],dp[v-a[i]][u-b[i]] + w[i])
            }
        }
    }
    return dp[V][U]
```

# 滚动数组-多重背包
```
    dp[0][0] = 0
    //二进制优化
    new_n := 0
    new_a := []
    new_b := []
    new_w := []
    for k := 1; k <= n[i](表示最多取的次数); k++ {
        new_n++
        new_a = a[i] * k
        new_b = b[i] * k
        new_w = w[i] * k
        n[i] -= k
    }
    if n[i] > 0 {
         new_n++
        new_a = a[i] * n[i]
        new_b = b[i] * n[i]
        new_w = w[i] * n[i]  
    } 
    
    //0-背包
    for i:=0; i < new_n; i ++ {
        for v :=V;v >=new_a[i];v-- {
            for u:=U; u >= new_b[i]; u-- {
                dp[v][u] =max(dp[u][v],dp[v - new_a[i]][u - new_b[i]] + w[i])
            }
        }
    
    }
    return dp[V][U]
```

# 滚动数组--混合背包 （略）

# demo
474 https://leetcode.cn/problems/ones-and-zeroes/description/


