package tools

import (
	"github.com/shopspring/decimal"
)

func YuanStringToCent(yuan string) int64 {
	yd, err := decimal.NewFromString(yuan)
	if err != nil {
		return 0
	}
	return YuanToCent(yd)
}

func YuanFloatToCent(yuan float64) int64 {
	return YuanToCent(decimal.NewFromFloat(yuan))
}

func YuanToCent(yuan decimal.Decimal) int64 {
	return yuan.Mul(decimal.NewFromInt(100)).IntPart()
}

func CentStrToYuan(cent string) float64 {
	cd, err := decimal.NewFromString(cent)
	if err != nil {
		return 0
	}
	return CentToYuan(cd)
}

func CentIntToYuan(cent int64) float64 {
	return CentToYuan(decimal.NewFromInt(cent))
}

func CentToYuan(cent decimal.Decimal) float64 {
	yuan, _ := cent.Div(decimal.NewFromInt(100)).Float64()
	return yuan
}

func YuanMulRatio(yuan float64, ratio int64) float64 {
	out, _ := decimal.NewFromFloat(yuan).Mul(decimal.NewFromInt(ratio)).Float64()
	return out
}
