package B_Tree
/**
	分裂：
		中间节点：和bTree一样
		叶子结点： bTree和双向链表结合
 */

/**
	分裂叶子节点
*/
func (b *BAddTree) splitLeaf (p *BAddTreeNode) *BAddTreeNode {
	//初始化right
	right := initBAddTreeNode(true)

	//获取当前节点个数
	sum := p.keyNum

	//获取上移节点
	mid := sum / 2 + 1

	//右边节点个数
	right.keyNum = M - mid

	//当前个数
	p.keyNum = mid

	//构造right.keys,right.children
	j := 1
	i := mid + 1
	for i <= M {
		right.keys[j] = p.keys[i]
		right.data[j] = p.data[i]
		i ++
		j ++
	}

	//处理mid上移
	midKey := p.keys[mid]
	PParent := p.p

	//mid上移
	if PParent == Nil {
		PParent = initBAddTreeNode(false)

		//修改指向
		p.p = PParent
		PParent.keys[1] = midKey
		PParent.children[0] = p
		PParent.children[1] = right
		right.p = PParent
		PParent.keyNum ++

	} else {
		//合并写法
		kIndex := PParent.keyNum + 1
		k := PParent.keyNum
		for k > 0 && PParent.keys[k] >= midKey {
			//key
			PParent.keys[kIndex] = PParent.keys[k]
			PParent.children[kIndex] = PParent.children[k]
			k--
			kIndex--
		}
		//循环完之后k的位置实际上是 target-1
		PParent.keys[k + 1] = midKey
		PParent.children[k] = p //可以不要
		PParent.children[k + 1] = right
		right.p = PParent
		PParent.keyNum ++
	}

	//修改p和right的链表指针
	Next := p.next
	right.next = Next
	if Next != Nil {
		Next.pre = right
	}
	p.next = right
	right.pre = p

	//返回父
	return PParent
}



/**
	分裂中间节点和b树一样
 */
func (b *BAddTree) splitMid (p *BAddTreeNode) *BAddTreeNode {
	//初始化right
	right := initBAddTreeNode(false)

	//获取当前的节点个数
	sum := p.keyNum

	//获取中间节点--中间节点上移
	mid := sum / 2 + 1

	//右边节点个数
	right.keyNum = M - mid
	p.keyNum = mid - 1

	//构造right.Keys & right.children
	j := 1
	i := mid + 1
	for ; i <= M; i ++ {
		right.keys[j] = p.keys[i]
		right.children[j - 1] = p.children[i - 1]
		right.children[j - 1].p = right
		j ++
	}
	right.children[j - 1] = p.children[i - 1]
	right.children[j - 1].p = right

	//处理mid
	midKey := p.keys[mid]
	p.keys[mid] = NULL

	//PParent
	PParent := p.p

	//mid上移
	if PParent == Nil {
		PParent = initBAddTreeNode(false)

		//修改指向
		p.p = PParent
		PParent.keys[1] = midKey
		PParent.children[0] = p
		PParent.children[1] = right
		right.p = PParent
		PParent.keyNum ++
		return PParent

	} else {
		//合并写法
		kIndex := PParent.keyNum + 1
		k := PParent.keyNum
		for k > 0 && PParent.keys[k] >= midKey {
			//key
			PParent.keys[kIndex] = PParent.keys[k]
			PParent.children[kIndex] = PParent.children[k]
			k--
			kIndex--
		}
		//循环完之后k的位置实际上是 target-1
		PParent.keys[k + 1] = midKey
		PParent.children[k] = p //可以不要
		PParent.children[k + 1] = right
		right.p = PParent
		PParent.keyNum ++

	}

	//返回
	return PParent
}



