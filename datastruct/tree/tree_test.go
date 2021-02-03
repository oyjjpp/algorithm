package tree

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	data := &Node{
		val: 1,
		left: &Node{
			val:  2,
			left: &Node{val: 4}},
		right: &Node{val: 3},
	}
	pushOne(data)
	printlnTree(data)
}

// printlnTree
// 前序遍历
func printlnTree(data *Node) {
	if data == nil {
		return
	}
	println(data.val)
	if data.left != nil {
		printlnTree(data.left)
	}

	if data.right != nil {
		printlnTree(data.right)
	}
}

func TestIsSameThee(t *testing.T) {
	data1 := &Node{
		val: 1,
		left: &Node{
			val:  2,
			left: &Node{val: 4}},
		right: &Node{val: 3},
	}
	data2 := &Node{
		val: 2,
		left: &Node{
			val:  2,
			left: &Node{val: 4}},
		right: &Node{val: 3},
	}
	rs := isSameThee(data1, data2)
	t.Log(rs)
}

func TestIsValidBST(t *testing.T) {
	data1 := &Node{
		val: 5,
		left: &Node{
			val:  2,
			left: &Node{val: 1},
			right: &Node{
				val:  4,
				left: &Node{val: 3},
			},
		},
		right: &Node{
			val:   6,
			right: &Node{val: 7},
		},
	}
	rs1 := isValidBST(data1)
	t.Log(rs1)

	data2 := &Node{
		val: 10,
		left: &Node{
			val: 5,
		},
		right: &Node{
			val:   15,
			left:  &Node{val: 6},
			right: &Node{val: 20},
		},
	}
	rs2 := isValidBST(data2)
	t.Log(rs2)
}

func TestIsInBST(t *testing.T) {
	data1 := &Node{
		val: 5,
		left: &Node{
			val:  2,
			left: &Node{val: 1},
			right: &Node{
				val:  4,
				left: &Node{val: 3},
			},
		},
		right: &Node{
			val:   6,
			right: &Node{val: 7},
		},
	}
	rs1 := isInBST(data1, 60)
	t.Log(rs1)
}

func TestInsertIntoBST(t *testing.T) {
	data1 := &Node{
		val: 5,
		left: &Node{
			val:  2,
			left: &Node{val: 1},
			right: &Node{
				val:  4,
				left: &Node{val: 3},
			},
		},
		right: &Node{
			val:   6,
			right: &Node{val: 7},
		},
	}
	rs1 := insertIntoBST(data1, 10)
	printlnTree(rs1)
}

func TestInvertBT(t *testing.T) {
	data := &Node{
		val: 4,
		left: &Node{
			val:   2,
			left:  &Node{val: 1},
			right: &Node{val: 3},
		},
		right: &Node{
			val:   7,
			left:  &Node{val: 6},
			right: &Node{val: 9},
		},
	}

	rs := invertBT(data)
	printlnTree(rs)
}
