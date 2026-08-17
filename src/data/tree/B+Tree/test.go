package B_Tree



func Test()  {
	//插入
	baTree := initBAddTree()
	baTree.insert(96)
	baTree.insert(40)
	baTree.insert(20)
	baTree.insert(39)
	baTree.insert(41)
	baTree.insert(42)
	baTree.insert(46)
	baTree.insert(43)
	baTree.insert(44)
	baTree.Print()

	baTree.Remove(96)
	baTree.Print()
	baTree.Remove(41)
	baTree.Print()
	baTree.Remove(46)
	baTree.Print()
	baTree.Remove(40)
	baTree.Print()
	baTree.Remove(43)
	baTree.Print()
	baTree.Remove(44)
	baTree.Print()
	baTree.Remove(42)
	baTree.Print()

}

