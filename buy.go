package main

import "math/big"

func Buy(_usdt int64, distribution uint8, probability float64) int64 {
	USDT += _usdt
	midToken := big.NewInt(0).Quo(AMMK, big.NewInt(USDT)).Int64()
	deltaToken := Token - midToken

	priceMove := GetRandomAccordingDistribution(distribution, probability) * 100
	earnToken := deltaToken * (100 + int64(priceMove)) / 100
	Token -= earnToken
	AMMK.Mul(big.NewInt(Token), big.NewInt(USDT))

	return earnToken
}
