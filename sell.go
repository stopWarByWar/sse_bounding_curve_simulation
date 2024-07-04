package main

import (
	"math/big"
	"math/rand"
)

// Sell
// @description 给定当前的价格，处理所有的holder
// @param holders 用户购买价格的sorted数组
// @param
// @return
func Sell(boughtUSDT int64, holders *[]int64, result *[]int64, earnBound, lossBound int64) bool {
	//挑出所有可能的发生价格变化的人的index的集合
	//随机挑出一个人执行
	//更新价格 and holder set
	res, newHolders := getActUser(*holders, boughtUSDT, earnBound, lossBound)
	if len(newHolders) == len(*holders) {
		return false
	}

	if res != 0 {
		*result = append(*result, res)
	}

	if USDT < InitUSDT {
		return true
	}
	return Sell(boughtUSDT, holders, result, earnBound, lossBound)
}

// getActUser
// @description 给定当前的价格，所有的holders，挑出所有可能的发生价格变化的人的index的集合 随机挑出一个执行的价格 以及更新的用户集合
// @param holders
// @return ，更新后holder集合
func getActUser(holders []int64, boughtUSDT int64, earnBound, lossBound int64) (int64, []int64) {
	var users []int64
	var idxs []int
	//todo: 不同的百分比出的概率不同
	for i, token := range holders {
		soldedUSDT := GetSellU(token)
		if (soldedUSDT-boughtUSDT)*100 > boughtUSDT*earnBound || (soldedUSDT-boughtUSDT)*100 < boughtUSDT*lossBound {
			users = append(users, token)
			idxs = append(idxs, i)
		}
	}

	if len(users) == 0 {
		return 0, holders
	}
	idx := idxs[rand.Intn(len(users))]
	soldU := SellDirectly(holders[idx])
	return soldU - boughtUSDT, append(holders[:idx], holders[idx+1:]...)
}

func GetSellU(trade int64) int64 {
	newToken := Token + trade
	newUSDT := big.NewInt(0).Quo(AMMK, big.NewInt(newToken)).Int64()
	return USDT - newUSDT
}

func SellDirectly(trade int64) int64 {
	Token += trade
	preU := USDT
	usdt := big.NewInt(0).Quo(AMMK, big.NewInt(Token))
	AMMK.Mul(big.NewInt(Token), usdt)
	USDT = usdt.Int64()
	return preU - USDT
}
