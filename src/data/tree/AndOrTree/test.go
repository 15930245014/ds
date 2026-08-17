package AndOrTree

import "fmt"

func Test()  {
	str := "(a | b &(c)) & d & (b | d & b)";
	answer := map[string]int{
		"a":1,
		"b":1,
		"c":-1,
		"d":-1,
	}
	tree,err := InitAndOrTree(str)
	if err != nil {
		fmt.Println(err.Error())
	}
	if tree != nil {
		FmtBoTree(tree.Root)
	}
	fmt.Println(DfsSolve(tree.Root,answer))
	fmt.Println(BFSSolve(tree.Root,answer))
}
