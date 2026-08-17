package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
	根据描述和分析图的最短路径完全满足动态规划的最优子结构性质
 */

 /**
 	图的表示方法采用邻接矩阵的表示方法
  */
/**

   采用邻接矩阵
*/
type AdjMatrix struct {
	vertex []string //顶点
	src [][]int  //邻接矩阵
	numVertex,numEdge int //顶点数目，边的数目
}
/**
	@author 许校
*注意 由于求最小的问题 故转化成负数来求解最大值问题
 */
func dynamicAdjMatrix (chart AdjMatrix) (int,string){
	//初始化转化负数
	for i := 0; i < chart.numVertex; i++ {
		for j := 0; j < chart.numVertex; j ++ {
			if chart.src[i][j] != 0 {
				chart.src[i][j] = -1 * chart.src[i][j]
			}
		}
	}

	//动态规划求解
	final := []int{}
	//初始化
	for i := 0; i < chart.numVertex; i ++ {
		final = append(final,0)
	}
	final[0] = 1

	result := map[int]int{}
	answer := map[int]string{}

	max := dynamicAdjM(chart,0,final,&result,answer)

	//求解最优解
	path := answer[max]

	//返回
	return -1 *max,path

}
/**
@param chart 图
@param begin 图的起点
@param final 顶点完成数组 避免回路
@param result 计算结果缓存机制
@param answer 计算路径数组结果

 */
func dynamicAdjM(chart AdjMatrix,begin int,final []int,result *map[int]int,answer map[int] string) int {
	//判断缓存
	if (*result)[begin] != 0 {
		return (*result)[begin]
	}


	//判断是否终止
	if begin == len(chart.src) - 1 {
		return 0
	}

	//定义负无穷
	max := math.MinInt64

	//寻找当前顶点的后记
	fg := false  //表示有后继
	for i := 1; i <len(chart.src); i ++ {
		//记录结果
		str := chart.vertex[begin]

		//判断
		if final[i] != 1 && chart.src[begin][i] != -1 * math.MaxInt64  {
			//表示有后继
			fg = true

			//深拷贝final
			/*final1 := make([]int,len(final))
			copy(final1,final)*/
			final[i] = 1

			//寻找子问题
			maxTmp := chart.src[begin][i] + dynamicAdjM(chart,i,final,result,answer)

			//判断最优质
			if maxTmp > max {
				//记录结果
				if answer[maxTmp - chart.src[begin][i]] != "" {
					str += "->"+answer[maxTmp - chart.src[begin][i]]
				} else {
					str += "->" + chart.vertex[i]
				}

				max = maxTmp
				answer[max] = str

			}
		}
	}

	//判断是否有后继
	if !fg {
		return 0
	}

	//写入缓存--表示begin 到终点的最短路径
	(*result)[begin] = max

	//返回
	return max

}

func main() {
	//构造邻接矩阵的图的算法省略
	chart := AdjMatrix{}

	//声明+无穷
	max :=  math.MaxInt64

	//构造顶点
	chart.numVertex = 9

	//顶点
	for i := 0; i < chart.numVertex; i ++ {
		chart.vertex = append(chart.vertex,"v" + strconv.Itoa(i))
	}

	//添加边
	chart.numEdge = 30
	chart.src = append(chart.src,[]int{0,1,5,max,max,max,max,max,max})
	chart.src = append(chart.src,[]int{1,0,3,7,5,max,max,max,max})
	chart.src = append(chart.src,[]int{5,3,0,max,1,7,max,max,max})
	chart.src = append(chart.src,[]int{max,7,max,0,2,max,3,max,max})
	chart.src = append(chart.src,[]int{max,5,1,2,0,3,6,9,max})
	chart.src = append(chart.src,[]int{max,max,7,max,3,0,max,5,max})
	chart.src = append(chart.src,[]int{max,max,max,3,6,max,0,2,7})
	chart.src = append(chart.src,[]int{max,max,max,max,9,5,2,0,4})
	chart.src = append(chart.src,[]int{max,max,max,max,max,max,7,4,0})

	fmt.Println(dynamicAdjMatrix(chart))
}



