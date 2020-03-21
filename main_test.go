package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFloat(t *testing.T) {
	rs1 := float64(99)
	rs2 := float64(9)
	rs := rs1 / rs2 * 10
	t.Logf("%f", rs)
	t.Logf("%f", Bcdiv(rs1, rs2, 2))
	t.Log(Bcdiv(rs1*10, rs2, 2))
	t.Log(GetDiscount(rs1, rs2))
}

func GetDiscount(showPrice, basePrice float64) float64 {
	round := Bcdiv(showPrice*10, basePrice, 2)
	if round == 0 && showPrice > 0 && basePrice > 0 {
		return 0.1
	} else {
		return round
	}
}

// Bcdiv
// 2个任意精度的数字除法计算
func Bcdiv(x, y float64, scale int) float64 {
	if y == 0 {
		return 0
	}
	rs := x / y

	data := round(rs, scale)
	fmt.Println("data", data)
	inst, _ := strconv.ParseFloat(data, 64)
	return inst
}

// round
func round(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return n
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

// FormatFloat64
func FormatFloat64(data interface{}) (float64, error) {
	str, _ := FormatString(data)
	return strconv.ParseFloat(str, 64)
}

// FormatString
// 转换为字符串
func FormatString(data interface{}) (string, error) {
	switch data.(type) {
	case json.Number:
		i, err := data.(json.Number).Int64()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(i, 10), nil
	case float32, float64:
		fdata := reflect.ValueOf(data).Float()
		return strconv.FormatFloat(fdata, 'f', -1, 64), nil
	case int, int8, int16, int32, int64:
		idata := reflect.ValueOf(data).Int()
		return strconv.FormatInt(idata, 10), nil
	case uint, uint8, uint16, uint32, uint64:
		udata := reflect.ValueOf(data).Uint()
		return strconv.FormatUint(udata, 10), nil
	case string:
		return data.(string), nil
	}
	return "", errors.New("invalid value type")
}

func TestBool(t *testing.T) {
	str := `{
		"code": 0,
		"msg": "ok",
		"body": {
			"albumId": 15,
			"asset": true,
			"programAssetList": [
				{
					"pId": "50000045",
					"asset": false
				},
				{
					"pId": "50000046",
					"asset": false
				},
				{
					"pId": "50000047",
					"asset": false
				}
			]
		}
	}`

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		t.Error(err.Error())
	}
	var mapData map[string]interface{}
	var ok bool
	if mapData, ok = data["body"].(map[string]interface{}); !ok {
		t.Error("断言异常")
	}

	if asset, ok := mapData["asset"].(bool); !ok {
		t.Error("bool 断言异常")
	} else {
		t.Log(asset)
	}

}

func TestTime(t *testing.T) {
	dateTime := secToDuration(-1)
	t.Log(dateTime)
}

func secToDuration(second int64) string {
	result := `00'00"`
	if second > 0 {
		dateStr := strconv.FormatInt(second, 10) + "s"
		dateTime, err := time.ParseDuration(dateStr)
		if err != nil {
			return result
		}
		rs := dateTime.String()
		rs = strings.Replace(rs, "h", "'", 1)
		rs = strings.Replace(rs, "m", "'", 1)
		rs = strings.Replace(rs, "s", `"`, 1)
		return rs
	}
	return result
}
