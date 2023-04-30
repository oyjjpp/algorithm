package hot100

import (
	"encoding/json"
	"sort"
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

func TestRemoveNthFromEndV(t *testing.T) {
	data := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	rs := removeNthFromEndV(data, 2)

	for rs != nil {
		t.Log(rs)
		rs = rs.Next
	}
}

func TestReverse(t *testing.T) {
	data := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	rs := reverse(data)

	for rs != nil {
		t.Logf("%d", rs.Val)
		rs = rs.Next
	}
}

func TestReverseList(t *testing.T) {
	data := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	reverseList(data)
}

func TestIsPalindromeList(t *testing.T) {
	data := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	rs := isPalindromeList(data)
	t.Log(rs)
}

func TestSearchOffer53(t *testing.T) {
	// nums1 := []int{2, 7, 11, 15}
	// nums2 := []int{1, 10, 4, 11}

	nums1 := []int{12, 24, 8, 32}
	nums2 := []int{13, 25, 32, 11}

	data := advantageCountC(nums1, nums2)
	t.Log(data)
}
func advantageCountC(nums1 []int, nums2 []int) []int {
	n := len(nums1)

	key1 := make([]int, n)
	key2 := make([]int, n)

	for i := 1; i < n; i++ {
		key1[i] = i
		key2[i] = i
	}

	sort.Slice(key1, func(i, j int) bool {
		return nums1[key1[i]] < nums1[key1[j]]
	})

	sort.Slice(key2, func(i, j int) bool {
		return nums2[key2[i]] < nums2[key2[j]]
	})

	res := make([]int, n)
	left, right := 0, n-1

	for i := 0; i < n; i++ {
		if nums1[key1[i]] > nums2[key2[left]] {
			res[key2[left]] = nums1[key1[i]]
			left++
		} else {
			res[key2[right]] = nums1[key1[i]]
			right--
		}
	}
	return res
}

func TestScanTreeCount(t *testing.T) {
	data := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	obj := CodecInit()
	rs := obj.serialize(data)
	t.Log(rs)
}

func TestBuildTreeX(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	rs := buildTreeX(inorder, postorder)

	str, err := json.Marshal(rs)
	t.Log(err, string(str))
}
