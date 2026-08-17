package BTree

/**
	keys索引从1开始
	children索引从0开始
	只有p.keys[1]存在前缀p.children[0]
	比p.keys[i]大的children索引为i 一一对应的

	函数列表：
	分裂:
	插入：
 */

/**
	插入节点
 */
func (b *BTree) Insert(val int)  {
	Root := b.Root
	if Root.KeyNum == 0 {
		//生成新节点
		newNode := initBNode()
		newNode.KeyNum = 1
		newNode.Keys[1] = val
		b.Root = newNode
		return
	}
	//寻找是否存在
	_,i,p := b.Search(val)
	if i == -1 {//没有查到
		//寻找插入位置
		j := p.KeyNum
		for  j > 0 && p.Keys[j] >= val {
			p.Keys[j + 1] = p.Keys[j]
			p.Children[j + 1] = p.Children[j]
			j--
		}

		//写入当前
		p.Keys[j + 1] = val
		p.Children[j + 1] = Nil
		p.KeyNum ++

		//判断是否需要分裂
		if p.KeyNum == M {
			parent := b.split(p)

			//判断是否需要更新root
			if parent.P  == Nil {
				b.Root = parent
			}

		}

	}
}

/**
    把当前p分裂
 */
func (b *BTree) split(p *BTNode) *BTNode {
	//初始化right
	right := initBNode()

	//获取当前的节点个数
	sum := p.KeyNum

	//获取中间节点--中间节点上移
	mid := sum / 2 + 1

	//右边节点个数
	right.KeyNum = M - mid
	p.KeyNum = mid - 1

	//构造right.Keys & right.children
	j := 1
	i := mid + 1
	for ; i <= M; i ++ {
		right.Keys[j] = p.Keys[i]
		right.Children[j - 1] = p.Children[i - 1]
		if right.Children[j - 1] != Nil {
			right.Children[j - 1].P = right
		}
		j ++
	}
	right.Children[j - 1] = p.Children[i - 1]
	if right.Children[j - 1] != Nil {
		right.Children[j - 1].P = right
	}

	//处理mid
	midKey := p.Keys[mid]
	p.Keys[mid] = NULL

	//PParent
	PParent := p.P

	//mid上移
	if PParent == Nil {
		PParent = initBNode()

		//修改指向
		p.P = PParent
		PParent.Keys[1] = midKey
		PParent.Children[0] = p
		PParent.Children[1] = right
		right.P = PParent
		PParent.KeyNum ++
		return PParent

	} else {
		//获取位置
		/*i := 1
		for i <= PParent.KeyNum && PParent.Keys[i] < midKey {
			i++
		}
		//写入 keys
		for j := PParent.KeyNum + 1; j >=i; j -- {
			PParent.Keys[j] = PParent.Keys[j - 1]
		}
		PParent.Keys[i] = KeyNum

		//写入children
		for j := PParent.KeyNum + 1; j >= i; j -- {
			PParent.Children[j] = PParent.Children[j - 1]
		}
		PParent.Children[i] = right*/
		//合并写法
		kIndex := PParent.KeyNum + 1
		k := PParent.KeyNum
		for k > 0 && PParent.Keys[k] >= midKey {
			//key
			PParent.Keys[kIndex] = PParent.Keys[k]
			PParent.Children[kIndex] = PParent.Children[k]
			k--
			kIndex--
		}
		//循环完之后k的位置实际上是 target-1
		PParent.Keys[k + 1] = midKey
		PParent.Children[k] = p //可以不要
		PParent.Children[k + 1] = right
		right.P = PParent
		PParent.KeyNum ++

	}

	//判断PParent是否满
	if PParent.KeyNum == M {
		return b.split(PParent)
	} else {
		return PParent
	}

}

