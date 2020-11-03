package leetcode

import(
	"testing"
	"log"
)

func TestRecoverTree(t *testing.T){
	data := &TreeNode{
		Val:3,
		Left:&TreeNode{Val:1},
		Right:&TreeNode{Val:4,Left:&TreeNode{Val:2}},
	}
	recoverTree(data)
	
	// 中序遍历
	inorder(data)
}

func TestRecoverTreeV2(t *testing.T){
	data := &TreeNode{
		Val:7,
		Left:&TreeNode{Val:2,Left:&TreeNode{Val:1},Right:&TreeNode{Val:3}},
		Right:&TreeNode{Val:6,Left:&TreeNode{Val:5},Right:&TreeNode{Val:4}},
	}
	recoverTreeV2(data)
	
	// 中序遍历
	inorder(data)
}

// 中序遍历
// inorder
func inorder(root *TreeNode){
	if root==nil{
		return
	}
	inorder(root.Left)
	log.Println(root.Val)
	inorder(root.Right)
}

func TestInorder(t *testing.T){
	data := &TreeNode{
		Val:1,
		Left:&TreeNode{Val:2,Left:&TreeNode{Val:21}},
		Right:&TreeNode{Val:3},
	}
	inorder(data)
}
