package main

import (
	"fmt"
	"math"
)

/**
	van Emde boas 树(vEB):vEB 树存在的意义是突破 O(lgn) 操作性能的极限。在基于比较的动态集合数据结构中，是存在这样的极限的。如同基于比较的排序方法也有 O(nlgn) 的性能极限一样。
vEB 树只在有限大小的全域 u 中有效，采用位图思想。
提前说一个结论，完全体的 vEB 树所有操作都可以在 O(lglgu) 时间内完成，其中 u 是树的全域。
min 和 max 两个属性是减少递归次数的关键，书中分析了这两个属性的多种作用，能够降低多个操作的复杂性。可以说 vEB 树的性能关键就在于对 min 和 max 的巧妙利用。
其中 min 不存储在树中的任何一个簇中。

	vEB树的注意事项：
	1、求解最大和最小值的操作不需要递归直接返回根的max和min 即可
	2、successor（后继）操作可以避免一个用于判断值x的后继是否位于high(x)中的递归调用。x 如果存在簇中当且仅当x < x簇的max
	3、通过min和max可以在O(1)时间判断vEB树是否为空、含有一个元素、或2个以上；如果min.max均为nil vEB树为空 ,min,max 均不为空且相等则vEB树仅含一个元素，min,max 均不为空且不相等则vEB树包含2个以上元素
   4、如果一颗vEB树为空。可以更新它的min 和 max值来实现插入一个元素。如果一棵vEB树仅含一个元素也可仅更新min和max来删除元素。
 */
const NIL = -100 //定义nil用于标记min和max 为空

/**
	vEB 树结构体
 */
type vEBoasTree struct {
	min, max int //最小、最大值
	u        int //全域大小
	summary  *vEBoasTree
	cluster  []*vEBoasTree //簇
}

/**
	获取簇索引⬇
 */
func high(x int, v *vEBoasTree) int {
	return int(float64(x) / math.Floor(math.Sqrt(float64(v.u))))
}

/**
	获取在当前簇的位置⬇
 */
func low(x int, v *vEBoasTree) int {
	return x % int(math.Floor(math.Sqrt(float64(v.u))))
}

/**
	获取生成簇时的值⬆
 */
func index(x, y int, v *vEBoasTree) int {
	return int(math.Floor(math.Sqrt(float64(v.u))))*x + y
}

/**
	构建一个vEBTree
 */
func vEBTreeCreate(v *vEBoasTree, u int) {
	//判断u
	if u <= 2 {
		v.min, v.max = NIL, NIL
		v.u = u
	} else {
		//初始化
		v.min, v.max = NIL, NIL
		v.u = u

		//初始化cluster
		su := int(math.Floor(math.Sqrt(float64(u))))
		for i := 0; i < su; i ++ {
			//放入
			v.cluster = append(v.cluster, new(vEBoasTree))

			//递归修改
			vEBTreeCreate(v.cluster[i], su)
		}

		//初始化summary
		v.summary = new(vEBoasTree)

		//递归构建summary
		vEBTreeCreate(v.summary, su)

	}
}

/**
	获取最大值
 */
func vEBTreeMax(v *vEBoasTree) int {
	return v.max
}

/**
	获取最小值
 */
func vEBTreeMin(v *vEBoasTree) int {
	return v.min
}

/**
  判断一个元素是否在vEBoasTree中
 */
func vEBoasTreeSearch(v *vEBoasTree, x int) bool {
	//判断
	if v.min == x || v.max == x {
		return true
	} else if v.u == 2 {
		return false
	} else {
		return vEBoasTreeSearch(v.cluster[high(x, v)], high(x, v))
	}
}

/**
	查找后继
 */
func vEBTreeSuccessor(v *vEBoasTree, x int) int {
	//判断是否为基本情况
	if v.u == 2 {
		if x == 0 && v.max == 1 {
			return 1
		} else {
			return NIL
		}

	} else if v.min != NIL && x < v.min { //是否为其它簇
		return v.min

	} else {
		//判断是否在当前的簇
		clusterMax := vEBTreeMax(v.cluster[high(x, v)])

		//判断当前簇是否存在后继
		if clusterMax != NIL && low(x, v) < clusterMax { //当前簇中存在
			//查找当前簇
			offSet := vEBTreeSuccessor(v.cluster[high(x, v)], low(x, v))

			//返回值
			return index(high(x, v), offSet, v)

		} else { //x的后继在其他簇中
			nextCluster := vEBTreeSuccessor(v.summary, high(x, v))

			//判断簇
			if nextCluster == NIL {
				return NIL
			} else {
				//获取最小值
				offSet := vEBTreeMin(v.cluster[nextCluster])

				//返回
				return index(nextCluster, offSet, v)
			}
		}

	}

}

/**
	获取前驱
 */
 func vEBTreePredecessor(v * vEBoasTree,x int) int {
 	//判断基本境况
 	if v.u == 2 {
 		if x == 1 && v.min == 0 {
 			return 0
		}  else {
			return NIL
		}

	} else if v.max != NIL && x > v.max {
		//前驱簇的最大值
		return v.max

	} else {
		//获取当前簇的最小值
		vEBmin := vEBTreeMin(v.cluster[high(x,v)])

		//判断
		if vEBmin != NIL && low(x,v) > vEBmin {
			//获取索引
			offSet := vEBTreePredecessor(v.cluster[high(x,v)],low(x,v))

			//返回
			return index(high(x,v),offSet,v)

		} else {
			//查找前驱簇
			preCluster := vEBTreePredecessor(v.summary,high(x,v))

			//判断
			if preCluster == NIL {
				//判断当前是否为最小
				if v.min != NIL && x > v.min { //x前驱min只有一个值 则为NIL
					return v.min
				} else {
					return NIL
				}
			} else {
				//获取索引---前簇的最大值
				offSet := vEBTreeMax(v.cluster[preCluster])

				//返回
				return index(preCluster,offSet,v)

			}
		}
	}
 }
/**
	插入元素
 */
 func vEBTreeInsert(v *vEBoasTree,x int) {
 	//判断树是否为空
 	if v.min == NIL {
 		v.min = x
 		v.max = x
	} else {
		if x < v.min  {
			v.min,x = x,v.min
			//v.min = x;
		}
		//如果域大小>2
		if v.u > 2 {
			//判断簇是否为空
			if vEBTreeMin(v.cluster[high(x,v)]) == NIL {
				vEBTreeInsert(v.summary,high(x,v))
				v.cluster[high(x,v)].min,v.cluster[high(x,v)].max = low(x,v),low(x,v)
			} else {
				vEBTreeInsert(v.cluster[high(x,v)],low(x,v))
			}
		}
		//判断
		if x > v.max {
			v.max = x
		}
	}
 }

 /**
 	删除元素
  */
  func vEBTreeDelete(v *vEBoasTree,x int) {
	if v.min == v.max { //当前只有一个元素时
		v.min , v.max = NIL,NIL
	} else if v.u == 2 { //基本境况
		if x == 0 {
			v.min = 1
		} else {
			v.min = 0
		}
		v.max = v.min

	} else {//2个以上元素且u >=4
		/**
			1、先判断当前x 是否为最小元素
				那么将当前node下的最小簇的最小节点用来顶替min
				并且将x置为最小簇的最小节点
			2、 如果删掉之后，x所在的簇为null,那么删除summary中的对应节点，用现存最大节点顶替max
		    3、如果x所在簇不为null，但是x==max 那么用x所在簇的最大节点顶替max
		 */
		 if x == v.min {
		 	firstCluster := vEBTreeMin(v.summary)
		 	x  = index(firstCluster,vEBTreeMin(v.cluster[firstCluster]),v)
		 	v.min = x
		 }
		 //簇中删除x
		 vEBTreeDelete(v.cluster[high(x,v)],low(x,v))

		 //判断x所在的簇是否为空
		 if vEBTreeMin(v.cluster[high(x,v)]) == NIL {
		 	//summary删除簇
		 	vEBTreeDelete(v.summary,high(x,v))

		 	//判断x 是否为最大顶点
		 	if x == v.max {
		 		maxCluster := vEBTreeMax(v.summary)
		 		if maxCluster == NIL {
		 			v.max = v.min
				} else {
					v.max = index(maxCluster,vEBTreeMax(v.cluster[maxCluster]),v)
				}
			}

		 } else {
		 	if x == v.max {
		 		v.max = index(high(x,v),vEBTreeMax(v.cluster[high(x,v)]),v)
			}
		 }


	}
  }

func main() {
	vEB := new(vEBoasTree)

	vEBTreeCreate(vEB, 16)

	vEBTreeInsert(vEB,3)
	vEBTreeInsert(vEB,2)
	vEBTreeInsert(vEB,4)
	vEBTreeInsert(vEB,5)

	fmt.Println(vEBTreePredecessor(vEB,4))
	vEBTreeDelete(vEB,3)
	fmt.Println(vEBTreePredecessor(vEB,4))
}
