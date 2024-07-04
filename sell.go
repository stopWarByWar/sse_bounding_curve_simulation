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
func Sell(usdt, token float64, holders []float64, currentPrice float64, result *[]float64, curve uint8, earnBound, lossBound float64) ([]float64, bool, float64) {
	//挑出所有可能的发生价格变化的人的index的集合
	//随机挑出一个人执行
	//更新价格 and holder set
	price, newHolders := getActUser(holders, currentPrice, earnBound, lossBound)
	if len(newHolders) == len(holders) {
		return newHolders, false, price
	}

	*result = append(*result, currentPrice-price)
	newPrice, _usdt, _token := updatePoolAfterSell(currentPrice, usdt, token, curve)
	if _usdt < InitUSDT {
		return nil, true, newPrice
	}
	return Sell(_usdt, _token, newHolders, newPrice, result, curve, earnBound, lossBound)
}

// getActUser
// @description 给定当前的价格，所有的holders，挑出所有可能的发生价格变化的人的index的集合 随机挑出一个执行的价格 以及更新的用户集合
// @param holders 用户购买价格的sorted数组
// @param currentPrice 当前curve价格
// @return 触发的sell的原始购买价格，更新后holder集合
func getActUser(holders []float64, currentPrice float64, earnBound, lossBound float64) (float64, []float64) {
	var users []float64
	var idxs []int
	//todo: 不同的百分比出的概率不同
	for i, buyPrice := range holders {
		priceChange := currentPrice - buyPrice
		if priceChange > buyPrice*earnBound || priceChange < buyPrice*lossBound {
			users = append(users, buyPrice)
			idxs = append(idxs, i)
		}
	}

	if len(users) == 0 {
		return currentPrice, holders
	}
	idx := idxs[rand.Intn(len(users))]
	price := holders[idx]
	return price, append(holders[:idx], holders[idx+1:]...)
}

func updatePoolAfterSell(currentPrice, usdt, token float64, curve uint8) (float64, float64, float64) {
	//todo: 卖的时候也加上随机
	token++
	newPrice, _ := Curve(curve, big.NewFloat(token)).Float64()
	//todo: 当前price为线性拟合，后续为了更高的精度可以用更精确的拟合
	usdt -= (currentPrice + newPrice) / 2
	return newPrice, usdt, token
}

func DirectSell(usdt, token, currentPrice float64, curve uint8) (float64, float64, float64) {
	return updatePoolAfterSell(currentPrice, usdt, token, curve)
}
