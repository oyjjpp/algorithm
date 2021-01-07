package leetcode

import (
	"strconv"
	"testing"
	"time"
)

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

func TestFindMedianSortedArrays(t *testing.T) {
	num2 := []int{2}
	num1 := []int{1, 3}
	rs := findMedianSortedArraysV2(num1, num2)
	if rs != 2 {
		t.Errorf("expect %f, result %f\n", 2.0, rs)
	}

	num1 = []int{1, 2}
	num2 = []int{3, 4}
	rs = findMedianSortedArraysV2(num1, num2)
	if rs != 2.5 {
		t.Errorf("expect %f, result %f\n", 2.5, rs)
	}
}

func TestLongestPalindrome(t *testing.T) {
	param := map[string]string{
		"babab": "babab",
		"babad": "bab",
		"cbbd":  "bb",
	}

	for k, v := range param {
		rs := longestPalindromeV2(k)
		if rs != v {
			t.Errorf("expect %s, result %s\n", v, rs)
		}
	}
}
