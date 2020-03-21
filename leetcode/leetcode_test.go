package leetcode

import (
	"strconv"
	"testing"
	"time"
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
	sumNode := addTwoNumbersV2(l1, l2)
	show(sumNode)
}

func TestDiv(t *testing.T) {
	curTime := Time() - 398
	rs := FormatTime(curTime)
	t.Log(rs)
}

// FormatTime
// 统一格式化时间
func FormatTime(dateTime int64) string {
	nowTime := Time()
	if dateTime != 0 {
		var desc string
		switch value := nowTime - dateTime; {
		case value > 0 && value <= 60:
			desc = "刚刚"
		case value > 60 && value <= 3600:
			diffTime := value / 60
			desc = strconv.FormatInt(diffTime, 10) + "分钟前"
		case value > 3600 && value <= 86400:
			diffTime := value / 3600
			desc = strconv.FormatInt(diffTime, 10) + "小时前"
		case value > 86400 && value <= 3*86400:
			diffTime := value / 86400
			desc = strconv.FormatInt(diffTime, 10) + "天前"
		case value > 3*86400:
			if Date("2006", nowTime) == Date("2006", dateTime) {
				// 同一年份
				desc = Date("1-02", dateTime)
			} else {
				desc = Date("2006-1-02", dateTime)
			}
		}
		return desc
	}
	// 默认返回当前时间
	return Date("2006-01-02", nowTime)
}

func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}

// Time
// 返回当前的 Unix 时间戳
func Time() int64 {
	timeLoc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(timeLoc).Unix()
}

func TestLengthOfLongestSubstring(t *testing.T) {
	param := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"dedf":     3,
		" ":        1,
	}
	for k, v := range param {
		rs := lengthOfLongestSubstring(k)

		if rs != v {
			t.Errorf("expect %d, result %d\n", v, rs)
		}
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	num1 := []int{1, 3}
	num2 := []int{2}
	rs := findMedianSortedArrays(num1, num2)
	if rs != 2 {
		t.Errorf("expect %f, result %f\n", 2.0, rs)
	}

	num1 = []int{1, 2}
	num2 = []int{3, 4}
	rs = findMedianSortedArrays(num1, num2)
	if rs != 2.5 {
		t.Errorf("expect %f, result %f\n", 2.5, rs)
	}
}
