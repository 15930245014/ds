# 混合背包
有 N 种物品和一个容量是 V 的背包。物品一共有三类：
第一类物品只能用1次（01背包）；
第二类物品可以用无限次（完全背包）；
第三类物品最多只能用 s[i] 次（多重背包）；
每种体积是 c[i]，价值是 w[i]。
求解将哪些物品装入背包，可使物品体积总和不超过背包容量，且价值总和最大。输出最大价值。
- s[i] == -1 表示 0-1背包
- s[i] == 0 表示完全背包
- s[i] > 0 表示多重背包

# 分组讨论 
分组讨论-伪代码
```
    for i := 0; i < N; i ++ {
        if s[i] == -1 {
            //0-1背包
            for v:=V; v >= c[i];v-- {
                dp[v] = max(dp[v],dp[v - c[i]] + w[i])  
            }
        } else if s[i] == 0 {
            //完全背包
            for v := c[i];v <=V; v ++ {
                dp[v] = max(dp[v],dp[v-c[i]] + w[i])
            }
        } else {
            //多重背包
            if n[i] * c[i] >= V {
                //完全背包
                for v := c[i];v <=V;v++ {
                    dp[v] = max(dp[v],dp[v - c[i]] + w[i])
                   
                }
                continue
            }
            //二进制优化
            new_c := []
            new_w := []
            new_n := 0
            for k:=1; k <= n[i]; k << 1 {
                new_c = append(new_c,k *c[i] )
                new_w = append(new_w,k *w[i] )
                new_n++
                n[i] -= k
            }
            if n[i] > 0 {
                new_c = append(new_c,n[i] *c[i] )
                new_w = append(new_w,n[i] *w[i] )
                new_n++
            }
            for m := 0; m < new_n;m++ {
                for v := V; v >= new_c[m];v-- {
                    dp[v] = max(dp[v],dp[v-new_c[m]] + w[m])
                }
            }
        }
        return dp[V]
    
    
    }
```
