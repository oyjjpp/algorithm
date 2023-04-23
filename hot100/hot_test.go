package hot100

import (
	"testing"
)

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

func TestPartition(t *testing.T) {
	root := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 12,
				Next: &ListNode{
					Val:  6,
					Next: &ListNode{Val: 9},
				},
			},
		},
	}
	scanList(root)
	t.Log("")
	rs := partition(root, 7)
	scanList(rs)
}

func TestRemoveElement(t *testing.T) {
	removeElement([]int{2, 3, 3, 2}, 3)
}

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	rs := isPalindrome(s)
	t.Log(rs)
}

func TestMaxDepth(t *testing.T) {
	data := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	rs := maxDepth(data)
	t.Log(rs)
}

func TestPrintBinaryLevel(t *testing.T) {
	data := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	printBinaryLevel(data)
}

func TestSubsets(t *testing.T) {
	data := []int{1, 2, 3}
	rs := subsets(data)
	t.Log(rs)
}

func TestCombine(t *testing.T) {
	rs := combine(4, 2)
	for item, key := range rs {
		t.Log(item, key)
	}

}

func TestPermuteRepeat(t *testing.T) {
	data := []int{1, 2, 3}
	rs := permuteRepeat(data)
	for _, value := range rs {
		t.Log(value)
	}
}

func TestMinDepth(t *testing.T) {
	data := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Left: &TreeNode{Val: 2},
	}

	rs := minDepth(data)
	t.Log(rs)
}

func TestSearchRange(t *testing.T) {
	rs := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	t.Log(rs)
}

func TestCommon(t *testing.T) {
	Common()
}

func TestMinWindowV(t *testing.T) {
	s := "ADOBECODEBANC"
	str := "ABC"
	rs := minWindowV(s, str)
	t.Log(rs)
}

func TestCheckInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidboaoo"
	rs := checkInclusion(s1, s2)
	t.Log(rs)
}

func TestRob_v2(t *testing.T) {
	temp := []int{1, 2, 3, 1}
	rs := rob_v2(temp)
	t.Log(rs)
}
