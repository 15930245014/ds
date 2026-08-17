 package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
-------图的定义
	1、图按照有无方向分为有向图和无向图。无向图由顶点和边组成，有向图由顶点和弧构成，弧有弧头和弧尾之分
	2、图按照弧或边的数目多少分为稀疏图和稠密图。任意两个顶点之间都存在边的叫无向完全图（n*(n-1)/2边），有向的叫做有向完全图（n*(n-1)弧）。若无重复的边或顶点到自身的边叫做简单图
	3、图中顶点之间有邻接点和依附的概念。无向图顶点的边数叫做度，有向图顶点分为入度和出度。
	4、图上的边或弧带权成为网
	5、图中顶点间存在路径，路径是指顶点到另一顶点间经过边或弧的数目，2顶点间存在路径说明是连通的，如果路径回到自身顶点的成为环。
	6、路径不重复经过一个顶点的叫做简单路径。若任意2个顶点是连通的则该图为连通图，有向的叫做强连通图。最大连通叫做连通分量，有向叫做强连通分量
 */

 /**
 	图的存储类型------邻接矩阵（Adjacency Matrix）
 	方式：用两个数组来表示图。-个一维数组储存图中顶点信息，一个二维数组（邻接矩阵）储存边或弧的信息。
 	举例：
 	顶点数组: vertex[4] = {A,B,C,D} ,邻接矩阵 arc[4][4] ==>src[i][j] 表示vertex[i] 顶点 和 vertex[j] 顶点是否邻接
 	src[i][j] = 1 表示存在线或弧 src[i][j] = 0 表示不连通
 	如果存在权值 则有src[i][j]不存在的值代表不连通（不一定为0） 连通 则src[i][j] 值为权

    根据邻接矩阵求度
 	无向图： 求B的度 则 循环 for src[1] : 判断src[1][j] 边存在之和
 	有向图：求B的度 则为 入度+出度  求入读：src[n][1] + 求出度 src[1][n]
  */

  //定义图的最大顶点数目
  const M = 40

  //构造邻接矩阵
  type AdjacencyMatrix struct {
  	 vertex [M]string
  	 src [M][M]int
  	 numVertex,numEdges int //顶点数目和边的数目
  }

  func main() {

  	//创建读取器
  	inputReader := bufio.NewReader(os.Stdin)

  	fmt.Printf("请输入图顶点的数目: ")

  	inputNumVertex, err := inputReader.ReadString('\n')

  	//判断
  	if err != nil {
  		fmt.Println("输入错误!")
	}
	//去掉\n
  	inputNumVertex = strings.Replace(inputNumVertex, "\n", "", -1)

  	//获取图顶点数目
  	numVertex,_ := strconv.Atoi(inputNumVertex)


  	//初始化图
  	adjacencyMatix := AdjacencyMatrix{}

  	//初始化邻接矩阵
  	adjacencyMatix.numVertex = numVertex //顶点数目

  	//初始化边
  	for i:= 0; i< numVertex; i ++ {
  		for j := 0; j < numVertex; j++ {
  			adjacencyMatix.src[i][j] = 0
		}

	}

  	//输入顶点
  	for i := 1; i <= numVertex; i ++ {
  		fmt.Printf("请输入第%d个顶点:",i)
  		vertex, err := inputReader.ReadString('\n')
  		if err != nil {
  			i--
  			fmt.Println("输入错误请重新输入")
  			continue
		}

  		//获取
  		adjacencyMatix.vertex[i - 1] = strings.Replace(vertex,"\n","",-1)
	}

  	//输入边
	  fmt.Printf("请输入图边或弧的数目: ")

	  inputNumEdgs, err := inputReader.ReadString('\n')

	  //判断
	  if err != nil {
		  fmt.Println("输入错误!")
	  }

	  inputNumEdgs = strings.Replace(inputNumEdgs,"\n","",-1)

	  //转化成整形
	  numEdgs,_ := strconv.Atoi(inputNumEdgs)

	  adjacencyMatix.numEdges = numEdgs

	  //输入边
	  for i := 1; i <= numEdgs; i ++ {
		  fmt.Printf("请输入第%d个边或弧 i,j,值形式:",i)
		  edge, err := inputReader.ReadString('\n')
		  if err != nil {
			  i--
			  fmt.Println("输入错误请重新输入")
			  continue
		  }

		  //获取
		  edge = strings.Replace(edge,"\n","",-1)

		  //分割成数组
		  edgeArr := strings.Split(edge,",")

		  //转化为整形
		  indexI,_ := strconv.Atoi(edgeArr[0])
		  indexj,_ := strconv.Atoi(edgeArr[1])
		  indexV,_ := strconv.Atoi(edgeArr[2])

		  //保存
		  adjacencyMatix.src[indexI][indexj] = indexV

	  }


	  //创建成功
	  fmt.Println("创建成功!")


  }
