package leetcode

import(
    "testing"
	"log"
    "encoding/json"
)

func TestInvertTree(t *testing.T){
    data := &TreeNode{
        Val:4,
        Left:&TreeNode{
            Val:2,
            Left:&TreeNode{Val:1},
            Right:&TreeNode{Val:3},
        },
        Right:&TreeNode{
            Val:7,
            Left:&TreeNode{Val:7},
            Right:&TreeNode{Val:9},
        },
    }
    rs := invertTree(data)
    temp, _ := json.Marshal(rs)
    log.Println(string(temp))
}
