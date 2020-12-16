package winning

import (
    "testing"
)

func TestProbability(t *testing.T){
    config := `{"A":5, "B":2, "C":2, "D":1}`
    for i := 0; i< 100; i++{
        rs := randomPobabilityV2(config)
        t.Log(rs)
    }
}
