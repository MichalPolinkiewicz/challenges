package main

import (
	"fmt"
	"math"
)

///https://www.codewars.com/kata/554a44516729e4d80b000012/train/go

func main() {
	fmt.Println(calculate(18000, 32000, 1500, 1.25))
}

func calculate(startPriceOld int, startPriceNew int, savingPerMonth int, percentLossByMonth float64) [2]int {
	months := 0
	var savedMoney float64
	percentLossByMonth = percentLossByMonth * 0.01

	oldCarPriceAsFloat := float64(startPriceOld)
	newCarPriceAsFloat := float64(startPriceNew)
	savingAsFloat := float64(savingPerMonth)

	calcNewValue := func(price, percentage float64) float64 {
		return price - price*percentage
	}

	for {
		if newCarPriceAsFloat >= savedMoney+oldCarPriceAsFloat {
			oldCarPriceAsFloat = calcNewValue(oldCarPriceAsFloat, percentLossByMonth)
			newCarPriceAsFloat = calcNewValue(newCarPriceAsFloat, percentLossByMonth)
			if months%2 == 0 {
				percentLossByMonth += 0.005
			}
			months++
			savedMoney += savingAsFloat
		} else {
			break
		}
	}
	return [2]int{months, int(math.Round(oldCarPriceAsFloat + savedMoney - newCarPriceAsFloat))}
}
