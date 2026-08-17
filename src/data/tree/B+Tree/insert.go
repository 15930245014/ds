package B_Tree

/**
	插入数据
 */
func (b *BAddTree) insert(val int) {
	Root := b.Root
	if Root.keyNum == 0 {
		//生成新节点
		newNode := initBAddTreeNode(true)
		newNode.keyNum = 1
		newNode.keys[1] = val
		newNode.data[1] = val
		b.Root = newNode
		return
	}
	//查找数据
	_,i,p,_:= b.search(val)
	if i == -1 {//没有查到
		//寻找插入位置
		j := p.keyNum
		for  j > 0 && p.keys[j] >= val {
			p.keys[j + 1] = p.keys[j]
			p.children[j + 1] = p.children[j] //可以忽略 叶子结点不存在
			j--
		}

		//写入当前
		p.keys[j + 1] = val
		p.children[j + 1] = Nil
		p.keyNum ++

		//判断是否需要分裂
		if p.keyNum == M {
			parent := b.split(p)
			//判断是否需要更新root
			if parent.p  == Nil {
				b.Root = parent
			}
		}
	}
}

/**
	分裂：叶节点 和 非叶节点逻辑不同
 */
func (b *BAddTree) split (p *BAddTreeNode) *BAddTreeNode {
	//获取leaf
	leaf := p.leaf

	//判断叶子结点
	var pParent * BAddTreeNode

	if leaf {
		pParent = b.splitLeaf(p)
	} else {
		pParent = b.splitMid(p)
	}

	//判断是否继续分裂
	if pParent.keyNum == M {
		return b.split(pParent)
	} else {
		return pParent
	}
}


