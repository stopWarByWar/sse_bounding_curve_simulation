package main

import (
	"math/big"
)

func Buy(currentPrice float64, curve uint8, usdt, token float64, distribution uint8, probability float64) (float64, float64, float64, float64) {
	token--
	realPrice := (GetRandomAccordingDistribution(distribution, probability) + 1) * currentPrice
	usdt += realPrice
	newPrice, _ := Curve(curve, big.NewFloat(token)).Float64()
	return realPrice, newPrice, usdt, token
}
