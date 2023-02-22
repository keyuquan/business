package utils

import (
	"fmt"
	"math"
	"time"
)

type TermSchema struct {
	Term    int64
	TermStr string
}

func SettleLotteryTermInt() *TermSchema {
	res := BetLotteryTerm()
	now := time.Now()
	if res.Term <= 1 {
		res.Term = 1440
		res.TermStr = now.Add(-1*time.Hour).Format("20060102") + "1440"
		return res
	}
	res.Term = res.Term - 1
	res.TermStr = now.Format("20060102") + fmt.Sprintf("%04d", res.Term)
	return res
}

func BetLotteryTerm() *TermSchema {
	res := new(TermSchema)
	now := time.Now()
	zero := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	res.Term = int64(math.Ceil(now.Sub(zero).Minutes()))
	res.TermStr = now.Format("20060102") + fmt.Sprintf("%04d", res.Term)
	return res
}
