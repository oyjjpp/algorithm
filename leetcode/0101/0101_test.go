package leetcode

import(
    "testing"
    "log"
)

func TestIsSymmetric(t *testing.T){
    data := &TreeNode{
        Val:1,
        Left:&TreeNode{
            Val:2,
            Left:&TreeNode{Val:3},
            Right:&TreeNode{Val:4},
        },
        Right:&TreeNode{
            Val:2,
            Left:&TreeNode{Val:4},
            Right:&TreeNode{Val:3},
        },
    }
    
    data = &TreeNode{
        Val:1,
        Left:&TreeNode{
            Val:2,
            Left:&TreeNode{Val:2},
        },
        Right:&TreeNode{
            Val:2,
            Right:&TreeNode{Val:2},
        },
    }
    
    // [1,2,2,2,null,2]
    rs := isSymmetric(data)
    log.Println(rs)
}
