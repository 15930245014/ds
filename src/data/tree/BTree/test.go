package BTree

func Test()  {
	//插入
	bsTree := InitBTree()
	bsTree.Insert(96)
	bsTree.Insert(40)
	bsTree.Insert(20)
	bsTree.Insert(39)
	bsTree.Insert(41)
	bsTree.Insert(42)
	bsTree.Insert(46)
	bsTree.Insert(43)
	bsTree.Insert(44)
	bsTree.Print()

	bsTree.Remove(42)

	bsTree.Print()
	bsTree.Remove(43)
	bsTree.Print()
	bsTree.Remove(46)
	bsTree.Print()
	bsTree.Remove(44)
	bsTree.Print()
	bsTree.Remove(96)
	bsTree.Print()
	bsTree.Remove(20)
	bsTree.Remove(40)
	bsTree.Remove(39)
	bsTree.Remove(41)
	bsTree.Print()
}
