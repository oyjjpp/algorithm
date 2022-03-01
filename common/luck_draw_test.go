package common

import "testing"

func BenchmarkAliasMethod(b *testing.B) {
	gifts := []Gift{
		{
			Name:        "mac",
			Probability: 1,
		},
		{
			Name:        "phone",
			Probability: 19,
		},
		{
			Name:        "优惠券",
			Probability: 80,
		},
	}
	for i := 0; i < b.N; i++ {
		_ = aliasMethod(gifts)
	}
}

func BenchmarkViolence(b *testing.B) {
	gifts := []Gift{
		{
			Name:        "mac",
			Probability: 1,
		},
		{
			Name:        "phone",
			Probability: 19,
		},
		{
			Name:        "优惠券",
			Probability: 80,
		},
	}
	for i := 0; i < b.N; i++ {
		_ = violence(gifts)
	}
}

func BenchmarkDispersed(b *testing.B) {
	gifts := []Gift{
		{
			Name:        "mac",
			Probability: 1,
		},
		{
			Name:        "phone",
			Probability: 19,
		},
		{
			Name:        "优惠券",
			Probability: 80,
		},
	}
	for i := 0; i < b.N; i++ {
		_ = dispersed(gifts)
	}
}
