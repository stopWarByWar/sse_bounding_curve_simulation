package main

import (
	"math/big"
)

const (
	Const = iota
	Line
	AMM
)

var AMMK *big.Float
var ConstBasicTokenPrice float64
var LineSlop float64

// Curve
// @description，输入当前已经从pool被换出来的token数量，得到当前token price
// @param model uint8 curve种类：0 -> y = a ; 1 -> y = k*x
// @param x uint64 已经被从pool中换出的token数量
// @return 当前token价格
func Curve(model uint8, x *big.Float) *big.Float {
	switch model {
	case Const:
		return big.NewFloat(ConstBasicTokenPrice)
	case Line:
		return big.NewFloat(0).Mul(x, big.NewFloat(LineSlop))
	case AMM:
		x.Mul(x, x)
		return big.NewFloat(0).Quo(AMMK, x)
	}
	return big.NewFloat(0)
}
