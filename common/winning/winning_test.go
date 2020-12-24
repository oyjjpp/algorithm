package winning

import (
	"testing"
)

func TestProbability(t *testing.T) {
	config := `{"A":5, "B":2, "C":2, "D":1}`
	for i := 0; i < 100; i++ {
		rs := randomPobabilityV2(config)
		t.Log(rs)
	}
}

func TestRoll(t *testing.T) {
	num := 0
	for i := 0; i < 10000; i++ {
		if rs := roolV2(0.1, 10); rs {
			num++
		}
	}
	t.Log(num)
}
