package main

import (
	"math"
	"math/rand"
)

// GetRandomAccordingDistribution
// @description 给定随机波动范围以及随机分布，给出一个随机后的买贵/便宜的百分比。这里我们精确到1%
// @param distribution分布 0 -> 均匀分布 1 -> 概率分布函数是三角形
// @param probability 随机盈亏百分比bound
// @return 随机盈亏百分比
func GetRandomAccordingDistribution(distribution uint8, probability float64) float64 {
	u := rand.Float64()
	switch distribution {
	case 0:
		return probability * u
	case 1:

		var x float64
		if u < 0.5 {
			x = -probability + math.Sqrt(2*probability*probability*u)
		} else {
			x = probability - math.Sqrt(2*probability*probability*(1-u))
		}
		// 将 x 转换为整数
		return x
	}
	return 0
}
