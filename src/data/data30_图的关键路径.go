package main

import "strconv"

/**
	AOE网：
		在AOV网的基础上，在一个表示工程的带权有向图中，用顶点表示事件，用有向边表示活动，用边上的权值表示活动的持续时间，这种有向图的边表示活动的网称之为AOE网

  AOV网与AOE网的区别：
	AOV网是顶点表示活动的网，他只描述活动之间的制约关系，而AOE网用边表示活动的网，边上的权值表示活动的持续时间

	关键路径:
		路径上各个活动的所持续的时间的之和成为路径长度，从原点到会点具有最大长度的路径叫做关键路径，在路径上的活动叫做关键活动

	参数定义:
	1、事件的最早发生时间etv:即顶点vk的最早发生时间
	2、事件的最晚发生时间ltv:即顶点vk的最晚发生时间，也就是每个顶点对应的事件最晚需要开始的时间，超出次时间将会延误整个工期
	3、活动的最早开工时间ete:即弧ak的最早发生时间
	4、活动的最晚开工时间lte :即弧aK的最晚发生时间，也就是不推迟工期的最晚开始时间
		---由1和2可以求得3和4，然后再根据ete[k]是否与lte[k]相等来判断ak是否关键活动
*/

/**
	采用邻接表的形式
 */
/**
	邻接表边结构体
 */
type edgeVertex struct {
	index int //顶点索引
	weight int //权
	next *edgeVertex //下一个指针域和
}

/**
	邻接表顶点结构体
 */
type vertexBox struct {
	in int //入度之和
	data string //顶点标识
	firstEdge *edgeVertex //指向的第一个边
}

/**
	邻接表
 */
type ljChart struct {
	vertexBox []vertexBox
	numVertex,numEdge int //边，顶点数目
}

/**
	获取关键路径的事件的最早开始时间
 */
 func getCriticalPathEtv(chart ljChart)([]int,[]int) {
 	//当前拓扑排序
 	var stack []int

 	//拓扑排序序列
 	var stack2 []int

 	//存放事件的最早开始时间
 	var etv []int

 	/**
 	初始化
 	 */
 	 //查询入度为0的顶点
 	 for i := 0; i < chart.numVertex; i ++ {
 	 	//判断入度为0进栈
 	 	if chart.vertexBox[i].in == 0 {
 	 		stack = append(stack,i)
		}

	 }

 	 //初始化事件的最早开始时间
 	 for i := 0; i < chart.numVertex; i++ {
 	 	etv = append(etv,0)
	 }

 	 //定义顶点的数目
 	 count := 0

 	 //拓扑排序，并求最早开始时间
 	 for len(stack) > 0 {
 	 	//出栈
 	 	top := stack[len(stack) - 1]
 	 	stack = stack[:len(stack) - 1]

 	 	//顶点数目+1
 	 	count ++
 	 	stack2 = append(stack2,top)

 	 	//遍历邻接点
 	 	for e := chart.vertexBox[top].firstEdge; e != nil; e = e.next {
 	 		//寻找
 	 		k := e.index
 	 		if etv[k] < etv[top] + e.weight {
 	 			etv[k] = etv[top] + e.weight
			}
 	 		//入度-1
			chart.vertexBox[k].in --

			//判断入度
 	 		if chart.vertexBox[k].in == 0 {
 	 			stack = append(stack,k)
			}
		}

	 }

 	 //判断数目
 	 if count != chart.numVertex {
 	 	//返回错误
	 }

 	 //返回
 	 return etv,stack
 }

 /**
 	求解关键路径
  */
  func getCriticalPath(chart ljChart)[]string {
  	//获取事件最早开始间
  	etv,stack := getCriticalPathEtv(chart)


  	//获取事件的最晚开始时间
  	var ltv []int

  	//初始化---取最后一个
  	for i := 0; i < chart.numVertex;i ++ {
  		ltv = append(ltv,etv[chart.numVertex - 1])
	}

  	//求最晚开始时间
  	for len(stack) > 0 {
  		//获取当前
  		top := stack[len(stack) - 1]
  		stack = stack[:len(stack) - 1]

  		//比较
  		for e := chart.vertexBox[top].firstEdge; e != nil ; e = e.next {
  			//获取k
  			k := e.index

  			//判断
  			if ltv[k] - e.weight < ltv[top] {
  				ltv[top] =  ltv[k] - e.weight
			}

		}
	}

  	//求解关键路径
  	var result []string

  	for i := 0; i < chart.numVertex; i ++ {
  		for e := chart.vertexBox[i].firstEdge; e != nil ; e = e.next {
  			k := e.index

  			//获取活动最早
  			ete := etv[i]

  			//活动最迟
  			lte := ltv[k] - e.weight

  			//判断是否相等
  			if ete == lte {
  				result = append(result,strconv.Itoa(i)+ strconv.Itoa(k))
			}

		}

	}

  	//返回
  	return result

  }