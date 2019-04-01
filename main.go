package main

import (
	"fmt"
	tree "learning-go/tree"
)

// t ->     2  (Leaf)
// t2 ->     10 (Tree)
//			    /	\
//		     2	2
//

func main() {
	t := tree.Tree{nil, 2, nil}
	t2 := tree.Tree{&t, 10, &t}
	fmt.Println(t.IsLeaf())
	fmt.Println(t2.IsLeaf())
	fmt.Println(t2.Sumatory())
	fmt.Println(t2.Ocurrence(10))
	fmt.Println(t2.Walk())
	fmt.Println(t2.NumberOfLeafs())
}
