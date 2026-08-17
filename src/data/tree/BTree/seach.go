package BTree

/**
	查询BTree:逐层查找
 */

func (b *BTree) Search(val int) (bool,int,*BTNode) {
	//判断根
	if b.Root == nil || b.Root.KeyNum == 0 {
		return false,-1,nil
	}

	//查询
	p := b.Root

	//查询到节点，如果没有查到返回当前节点
	var node *BTNode

	//遍历查询
	for p != Nil {
		//查询（可以二分，这里简单就直接轮询）
		i := 1
		for i<= p.KeyNum {
			if p.Keys[i] < val{
				i++
			} else if val == p.Keys[i] {
				return true,i,p
			} else {
				break
			}
		}

		//获取下一层
		if p.Children[i-1] == nil {
			node = p
		}
		p = p.Children[i-1]
	}

	//没有找到
	return false,-1,node
}
