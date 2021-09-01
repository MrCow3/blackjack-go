package main

import (
	"fmt"
	"time"
)

func dealer() {
	for {
		if dealerAmount < 21 {
			fmt.Println("The dealer hits")
			fmt.Println(dealerAmount)
			randCardDealer(0)
			time.Sleep(1 * time.Second)
		}

		if dealerAmount > 21 {
			fmt.Println("Dealer busted!")
			fmt.Println(dealerAmount)
			break
		}

		if dealerAmount == 21 {
			fmt.Println("Dealer Blackjack!")
			break
		}
	}
}
