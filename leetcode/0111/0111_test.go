package leetcode

import (
    "testing"
    "log"
)

// [3,9,20,null,null,15,7]
func TestMinDepthBFS(t *testing.T){
    data := &TreeNode{
        Val:3,
        Right:&TreeNode{Val:9},
        Left:&TreeNode{
            Val:20,
            Left:&TreeNode{Val:15},
            Right:&TreeNode{Val:7},
        },
    }
    minRes := minDepthBFS(data)
    log.Println(minRes)
}
