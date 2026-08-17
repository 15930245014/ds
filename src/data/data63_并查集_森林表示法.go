package main

/**
	并查集
	不相交集合森林：适用有根树来表示结合。树中每个结点包含一个成员，每棵树代表一个集合。在一个不相交集合森林中每个成员仅指向它的父结点。
	简单路径：find-set操作所访问的结点构成了查找简单路径

	改进运行时间的启发策略：
	1、按秩合并：让具有较小结点的树指向较多结点的树。
	2、路径压缩:压缩树的高度
 */

/**
	结合树数据结构
 */
type structureTree struct {
	p    *structureTree //指向父结点
	rank int            //高度
	val  byte           //数值
}

/**
	整个结合
 */
type sTreeArray struct {
	arr []*structureTree
}

/**
	构造结合
 */
func makeSTree(s sTreeArray, x byte)sTreeArray {
	//构造新结点
	nW := new(structureTree)
	nW.val = x
	nW.rank = 0
	nW.p = nW

	//填入结合
	s.arr = append(s.arr, nW)

	//返回
	return s
}

/**
	查询x为代表的树
 */
func findsTree(s sTreeArray, x byte) *structureTree {
	for _, value := range s.arr {
		if value.val == x {
			return value
		}
	}
	//返回
	return nil
}

/**
	findByNode---算法书中提及
 */
func findsTreeByNode(x *structureTree) *structureTree {
	//向上查找
	if x.p != x {
		x.p = findsTreeByNode(x.p)
	}
	//返回
	return x.p
}
/**
	合并--按秩合并
 */
func unionStree(s sTreeArray,x,y byte)sTreeArray {
	//查找结点
	xTree := findsTree(s,x)
	yTree := findsTree(s,y)

	//合并
	if xTree.rank > yTree.rank {
		yTree.p = xTree

		//集合中删除yTree--代码略

	} else {
		xTree.p = yTree

		//判断2者是否相等
		if xTree.rank == yTree.rank {
			yTree.rank ++
		}

		//集合中删除xTree---代码略

	}

	//返回
	return s

}
