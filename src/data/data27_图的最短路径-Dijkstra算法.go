package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
	最短路径，是指两顶点之间经过的边上的权值之和最少的路径(path)，路径的第一个顶点是源点最后一个顶点是终点。
	求解单元最短路径的常见算法：迪杰斯特拉（Dijkstra）、弗洛伊德（Floyd）,动态规划（dynamic programming）
 */

 /**target
 	迪杰斯特拉(Dijkstra)算法原理:求解原点与某个顶点的最短路径后，已此顶点为中心，再求解源点与此顶点邻接的顶点的最短路径（贪心策略）与之前的计算的源点到邻接顶点距离取最小，直到递推到源点。时间复杂度O(n^2)
  */


  /**
  	迪杰斯特拉算法实现思路：
  	声明final数组表示之前已经求得源点到各顶点的距离 如:final[w] = 1表示已经求得w与源点的距离。------避免重复（******相当于缓存机制）
    lowCost 数组表示源点到终点间(包括终点)的各顶点的最短路径  lowCost[w] 表示源点到顶点w的最短路径
  	parent 数组表示源点到顶点的前驱 源点到源点为0， 例如v0 -> v1 -> v2 此时p[2] = 1 表示 v0到到v2的必须通过v1

  	每次求得源点到其它顶点的最短路径（贪心策略），在以此顶点为中心求得源点到其他顶点的最短路径
  */

  /**
  	采用邻接矩阵
   */
   type adjMatrix struct {
   		vertex []string //顶点
   		src [][]int  //邻接矩阵
   		numVertex,numEdge int //顶点数目，边的数目
   }

   /**
   	迪杰斯特拉拉算法
    */
   func  Dijkstra(chart adjMatrix)(int,string){
   	/**
   		初始化
   	 */
   	 //定义final数组并初始化
   	 var final []int
   	 /*for i := 0; i < chart.numVertex; i++ {
   	 	final := append(final,0)
	 }*/

   	 //定义parent,并初始化
   	 var parent []int
   	 /*for i := 0; i < chart.numVertex; i++ {
   	 	parent := append(parent,0)
	 }*/

   	 //定义lowCost,并初始化--即开始源点到其它顶点的距离 不邻接的用+无穷表示
   	 var lowCost []int
   	 /*for i := 0; i < chart.numVertex;i ++ {
   	 	lowCost = append(lowCost,chart.src[0][i])
	 }*/

   	 //注释内容合并
   	 for i := 0; i < chart.numVertex; i ++ {
   	 	final = append(final,0)
   	 	parent = append(parent,0)
   	 	lowCost = append(lowCost,chart.src[0][i])
	 }

   	 //源点final = 1
   	 final[0] = 1
   	 lowCost[0] = 0

   	 //开始主循环，每次求得源点到某个顶点的最短路径 下标从1开始
   	 for i := 1; i < chart.numVertex; i ++ {
   	 	//定义最小距离,初始化为+无穷
   	 	min := math.MaxInt64

   	 	//源点到某个顶点的最短路径的顶点
   	 	target := -1

		 //寻找源点到后续的最近的顶点
		 for j := 1; j < len(lowCost); j ++ {
		 	//判断
		 	if final[j] != 1 && lowCost[j] < min { //final[j] 表示已经求得了
		 		min = lowCost[j]
		 		target = j
			}

		 }

		 //求得当前的最短路径
		 final[target] = 1


		 //以当前顶点为中心更新lowCost
		 for k := 1; k < chart.numVertex; k ++ {
		 	if final[k] != 1 && (chart.src[target][k] != math.MaxInt64 && min + chart.src[target][k] < lowCost[k])  { //上面判断略有不同math.MaxInt64 + 任意正整数 为负数
		 		lowCost[k] = min + chart.src[target][k]
		 		parent[k] = target //更新前驱
			}

		 }

	 }


   	 //最短路径
   	 minPath := lowCost[chart.numVertex - 1]

   	 //求路径
   	 node := chart.numVertex - 1
   	 path := chart.vertex[node]
   	 for parent[node] != 0 {
   	 	node = parent[node]
   	 	path += "->" + chart.vertex[node]
	 }

   	 path += "->" + chart.vertex[0]

   	 //返回
   	 return minPath,path

   }

   /**
   	测试Dijkstra
    */
   func main() {
   	//构造邻接矩阵的图的算法省略
   	chart := adjMatrix{}

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

   	//调用
   	minPath,path := Dijkstra(chart)

   	fmt.Println("最短路径为:",minPath)
   	fmt.Println("路径为:",path)


   }
