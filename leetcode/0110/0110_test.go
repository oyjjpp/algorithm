package leetcode

import(
    "testing"
    "log"
)

func TestAbs(t *testing.T){
    rs := abs(-10)
    log.Println(rs)
    
    rs = abs(20)
    log.Println(rs)
}

func TestIsBalanced(t *testing.T){
    data := &TreeNode{
        Val:3,
        Left:&TreeNode{
            Val:9,
        },
        Right:&TreeNode{
            Val:20,
            Left:&TreeNode{Val:15},
            Right:&TreeNode{Val:7},
        },
    }
    bo := isBalancedV2(data)
    log.Println(bo)
}
