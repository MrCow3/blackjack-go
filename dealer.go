package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var dealerAmount int = 0

var cards = [9]int{2, 3, 4, 5, 6, 7, 8, 9, 10}

func dealer() {
	fmt.Println(dealerAmount)
	for {
		if dealerAmount < 21 {
			if dealerAmount < 21 {
				fmt.Println("The dealer hits")
				fmt.Println(dealerAmount)
				randCard(0)
			}
		}

		if dealerAmount > 21 {
			fmt.Println("Dealer busted!")
			fmt.Println(dealerAmount)
			os.Exit(0)
		}

		if dealerAmount == 21 {
			fmt.Println("Dealer Blackjack!")
			os.Exit(0)
		}

	}
}

func randCard(n int) int {
	rand.Seed(time.Now().UnixNano())
	in := cards
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	dealerAmount = dealerAmount + pick
	return 0
}
