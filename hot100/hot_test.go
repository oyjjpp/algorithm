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
