package leetcode

import(
    "log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric
// 对称二叉树
func isSymmetric(root *TreeNode) bool {
    data := []int{}
    
    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode){
        if root == nil{
            return
        }
        inorder(root.Left)
        data = append(data, root.Val)
        inorder(root.Right)
    }
    inorder(root)
    dataLen := len(data)
    if dataLen ==0 {
        return true
    }
    
    num := dataLen/2
    if dataLen%2 == 0{
        return false
    }
    data1 := data[:num]
    data2 := data[num+1:]
    log.Println(data, data1, data2)
    
    for i:=0;i<len(data1);i++{
        log.Println(data1[i], data2[num-i-1])
        if(data1[i]!=data2[num-i-1]){
            return false
        }
    }
    return true 
}
