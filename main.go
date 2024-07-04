package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
)

const (
	InitUSDT        = 30_000_000_000
	InitToken       = 1_000_000_000_000_000_000
	TargetUSDT      = 115_000_000_000
	PurchaseUAmount = 500_000_000

	EarnBound = 100
	LoseBound = -50

	MaxRound = 10000

	StartSell             = 5
	SellDirectProbability = 0.2
	SellDirectEarn        = 20

	Distribution         = 1
	ProbabilityMoveBound = 0.5
)

func main() {
	success := 0
	failed := 0
	no_result := 0

	totalSuccesRound := 0
	toTalFailRound := 0

	for i := 0; i < 1000; i++ {
		round, result := test()
		if result {
			success++
			totalSuccesRound += round
		} else if round != MaxRound {
			failed++
			toTalFailRound += round
		} else {
			no_result++
		}
	}
	fmt.Printf("sucess:%d,avg round:%d\nfail:%d,avg round:%d\nnoresult:%d", success, totalSuccesRound/success, failed, toTalFailRound/failed, no_result)
}

func test() (int, bool) {
	var holders []int64
	var result []int64

	USDT = InitUSDT
	Token = InitToken
	AMMK = big.NewInt(0).Mul(big.NewInt(USDT), big.NewInt(Token))

	for i := 0; i < MaxRound; i++ {
		var stop bool
		_token := Buy(PurchaseUAmount, Distribution, ProbabilityMoveBound)
		if USDT > TargetUSDT {
			fmt.Printf("Sucessfull:%d, usdt: %d,token:%d,result:%d,holders:%d\n", i, USDT, Token, len(result), len(holders))
			return i, true
		}
		if i < StartSell {
			holders = append(holders, _token)
			continue
		} else {
			r := rand.Float64()
			earnedU := GetSellU(_token) - PurchaseUAmount
			if earnedU > SellDirectEarn*PurchaseUAmount/100 && r < SellDirectProbability {
				SellDirectly(_token)
				result = append(result, earnedU)
			} else {
				holders = append(holders, _token)
				stop = Sell(PurchaseUAmount, &holders, &result, EarnBound, LoseBound)
				if stop {

					fmt.Printf("Faield:%d,usdt: %d,token:%d,result:%v,holders:%v\n", i, USDT, Token, len(result), len(holders))
					return i, false
				}
			}
		}
		//fmt.Printf("round %d,usdt: %d,token:%d,result:%d,holders:%d\n", i, USDT, Token, len(result), len(holders))
	}
	fmt.Printf("No Result: usdt: %d,token:%d,result:%v,holders:%v\n", USDT, Token, len(result), len(holders))
	return MaxRound, false
}
