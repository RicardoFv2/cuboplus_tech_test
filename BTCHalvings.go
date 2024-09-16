package main

import (
	"fmt"
)

const (
	maxBitcoinSupply = 21_000_000
	satsPerBitcoin   = 100_000_000
	initialReward    = 50 * satsPerBitcoin
	halvingInterval  = 210_000
	totalHalvings    = 32
)

func main() {

	totalSupply := 0.0
	reward := initialReward
	totalMined := 0.0
	halvingCount := 0

	for halvingCount <= totalHalvings {

		minedBeforeHalving := float64(reward) / satsPerBitcoin * float64(halvingInterval)
		totalMined += minedBeforeHalving
		totalSupply = totalMined

		percentMined := (totalMined / maxBitcoinSupply) * 100

		fmt.Printf("Halving %d:\n", halvingCount)
		fmt.Printf("Block reward: %d SATS\n", reward)
		fmt.Printf("Total Bitcoin mined: %.2f BTC\n", totalSupply)
		fmt.Printf("Percentage of total supply mined: %.8f%%\n", percentMined)
		fmt.Println("------------------------")

		reward /= 2
		halvingCount++
	}
}
