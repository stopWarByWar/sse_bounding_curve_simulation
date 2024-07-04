package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
)

const (
	InitUSDT              = 30
	InitToken             = 1_000_000_000
	TargetUSDT            = 115
	EarnBound             = 1.0
	LoseBound             = -0.5
	MaxRound              = 10000
	StartSell             = 5
	SellDirectProbability = 0.1
	SellDirectEarn        = 0.1
	Distribution          = 1
	ProbabilityMoveBound  = 0.5
)

// const initUSDT =
func main() {
	var holders []float64
	var result []float64

	usdt := float64(InitUSDT)
	token := float64(InitToken)
	AMMK = big.NewFloat(usdt).Mul(big.NewFloat(usdt), big.NewFloat(token))
	var curve uint8 = AMM
	//ConstBasicTokenPrice =
	//LineSlop int64
	currentPrice, _ := Curve(AMM, big.NewFloat(token)).Float64()

	for i := 0; i < MaxRound; i++ {
		var buyPrice float64
		var stop bool
		buyPrice, currentPrice, usdt, token = Buy(currentPrice, curve, usdt, token, Distribution, ProbabilityMoveBound)
		if usdt > TargetUSDT {
			fmt.Printf("Sucessfull\nusdt: %f\ntoken:%f\nresult:%v\nholders:%v\n", usdt, token, result, holders)

		}
		holders = append(holders, buyPrice)
		if i < StartSell {
			continue
		} else {
			r := rand.Float64()
			if (currentPrice-buyPrice) > buyPrice*SellDirectEarn && r < SellDirectProbability {
				currentPrice, usdt, token = DirectSell(usdt, token, currentPrice, curve)
			} else {
				holders, stop, currentPrice = Sell(usdt, token, holders, currentPrice, &result, curve, EarnBound, LoseBound)
				if stop {
					fmt.Printf("Faield\nusdt: %f\ntoken:%f\nresult:%v\nholders:%v\n", usdt, token, result, holders)
					return
				}
			}
		}
	}
	fmt.Printf("No Result\nusdt: %f\ntoken:%f\nresult:%v\nholders:%v\n", usdt, token, result, holders)
}
