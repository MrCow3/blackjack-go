package main

import (
	"fmt"
	"os"
)

func dealer() {
	for {
		if dealerAmount < 21 {
			if dealerAmount < 21 {
				fmt.Println("The dealer hits")
				fmt.Println(dealerAmount)
				randCardDealer(0)
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
