package tree

// Tree struct
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// IsLeaf
func (t *Tree) IsLeaf() bool {
	return t.Left == nil && t.Right == nil
}

// Walk
func (t *Tree) Walk() []int {
	arr := []int{}
	arr = append(arr, t.Value)
	if t.Left != nil {
		arr = append(arr, t.Left.Walk()...)
	}
	if t.Right != nil {
		arr = append(arr, t.Right.Walk()...)
	}
	return arr
}

// Sumatory
func (t *Tree) Sumatory() int {
	values := t.Walk()
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

// Ocurrence
func (t *Tree) Ocurrence(v int) int {
	ocurrence := 0
	if t.Value == v {
		ocurrence += 1
	}
	if t.Left != nil {
		ocurrence += t.Left.Ocurrence(v)

	}
	if t.Right != nil {
		ocurrence += t.Right.Ocurrence(v)
	}
	return ocurrence
}

// NumberOfLeafs

func (t *Tree) NumberOfLeafs() int {
	numLeafs := 0
	if t.IsLeaf() {
		numLeafs += 1
	} else {
		numLeafs += t.Left.NumberOfLeafs() + t.Right.NumberOfLeafs()
	}
	return numLeafs
}
