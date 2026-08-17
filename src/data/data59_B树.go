package main

/**
	B树：是一种多路自平衡搜索树主要用于内存和外存之间的数据交互。---本此代码 下标从1开始存储
	*B树的阶 即当前树结点的最大容量 一般 定义一个m 则有 树结点的最大容量为 2*m - 1
	*B树的高度增加只能通过分裂来实现 一般以中间关键字来分裂  分裂成2部分 各有 m - 1 元素（2（m - 1） + 1 == 2*m -1）且当且仅当树结点满时才分裂，然后中间关键字插入到父结点中，并以此类推
	*B树结点有key 关键字数组和 children[]指针数组和一个是否为叶子结点的标识，和统计元素个数
	*B树的拆入是从根开始的
 */
//定义阶
const T = 4

/**
	定义B树结点结构体
 */
type BtreeNode struct {
	n        int          //元素个数
	key      []int        //元素 ---升序排列
	leaf     bool         //是否为叶子结点
	children []*BtreeNode //孩子数组与key的下标一致
}

/**
	B树
 */
type Btree struct {
	root *BtreeNode
}

/**
	构造新结点---主要是实例化2个切片
 */
func initBtreeNode() *BtreeNode {
	//测试
	a := new(BtreeNode)

	for i := 0; i <= 2*T-1; i++ {
		//初始化为0
		a.key = append(a.key, 0)

		//初始化孩子
		a.children = append(a.children, nil)

	}

	//返回新结点
	return a
}

/**
	分裂函数
	分裂条件x的第x孩子满了
	*@params x *Btree 当前的父结点
	*@params i int 分裂x 的 第i个孩子
 */
func splitBtreeNode(x *BtreeNode, i int) {
	//构造另一半
	z := initBtreeNode();

	//获取当前的孩子
	y := x.children[i]

	//根据中间关键子进行分裂---关键字位置索引为T

	//1、处理z
	z.leaf = y.leaf
	z.n = T - 1
	for i := T + 1; i <= y.n; i ++ {
		z.key[i-T] = y.key[i]
		z.children[i-T] = y.children[i]

		//修改y为初始值
		y.key[i] = 0
		y.children[i] = nil

	}

	//2、处理y
	y.n = T - 1
	midKey := y.key[T]
	y.key[T] = 0
	y.children[T] = nil

	//3、处理x
	x.n = x.n + 1

	//先处理i+1前的children 然后children[i + 1] 指向 z [i] 还是指向y
	for j := x.n; j > (i + 1); j -- {
		x.children[j] = x.children[j-1]
	}

	//在i位置插入key
	for j := x.n; j > i; j -- {
		x.key[j] = x.key[j-1]
	}
	x.key[i] = midKey
	x.children[i] = y
	x.children[i+1] = z
}

/**
	向x插入结点树未满时
 */
func insetBtreeNotFull(x *BtreeNode, k int) {
	//获取当前结点的元素个数
	i := x.n

	//判断当前是否为叶子结点
	if x.leaf {
		//直接插入到当前
		for i >= 1 && x.key[i] >= k {
			//右移
			x.key[i+1] = x.key[i]
			x.children[i+1] = x.children[i]
		}

		//插入k
		x.key[i+1] = k
		x.children[i+1] = nil
		x.n ++

	} else { //插入到叶子
		//获取位置
		for i >= 1 && x.key[i] >= k {
			i --
		}

		//要插入位置
		i = i + 1

		//判断当前的叶子结点
		if x.children[i+1] == nil {
			x.children[i+1 ] = initBtreeNode()
		}

		//判断是否满
		if x.children[i+1].n == (2*T - 1) {
			//分裂
			splitBtreeNode(x, i)

			//判断
			if k > x.key[i] {
				i++
			}

		}

		//递归插入
		insetBtreeNotFull(x.children[i], k)

	}
}

/**
	添加记录
 */
func insetBtree(r *Btree, key int) {
	//获取根
	root := r.root

	//判断根是否满
	if root.n == 2*T-1 {
		//新结点
		newNode := initBtreeNode()
		newNode.children[1] = root
		newNode.leaf = false

		splitBtreeNode(newNode, 1)
		r.root = newNode

		insetBtreeNotFull(newNode, key)

	} else {
		insetBtreeNotFull(root, key)
	}
}

/**
	搜索B树
 */
func searchBtree(r *BtreeNode, key int) bool {
   	//现在当前结点查找
	i := 1
	for i <= r.n && r.key[i] < key {
   		i ++
	}

	//判断
	if i <= r.n && r.key[i] == key {
		return true
	}

	//查找其子
	if r.children[i] == nil {
		return false
	}

	//查找子
	return searchBtree(r.children[i],key)

}

func main() {
	//B树的删除略
}
