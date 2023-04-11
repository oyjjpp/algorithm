package hot100

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	sumNode := addTwoNumbers(l1, l2)
	scanList(sumNode)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	rs := lengthOfLongestSubstring(s)
	t.Log(rs)
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	rs := findMedianSortedArrays(nums1, nums2)
	t.Log(rs)
}

func TestFibV(t *testing.T) {
	rs := fibV(4)
	t.Log(rs)
}

func TestFibV2(t *testing.T) {
	rs := fibV2(4)
	t.Log(rs)
}

func TestCoinChange(t *testing.T) {
	coins := []int{1}
	amount := 10000
	rs := coinChange(coins, amount)
	t.Log(rs)
}

func TestCoinChangeV(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	rs := coinChangeV(coins, amount)
	t.Log(rs)
}

func TestMaxPathSum(t *testing.T) {
	data := &TreeNode{
		Val: -3,
		// Left: &TreeNode{Val: -1},
		// Right: &TreeNode{Val: 3},
	}
	rs := maxPathSum(data)
	t.Log(rs)
}
