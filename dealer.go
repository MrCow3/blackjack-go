package main

import (
	"fmt"
	"os"
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
			fmt.Println("You win")
			os.Exit(0)
		}

		if dealerAmount == 21 {
			fmt.Println("Dealer Blackjack!")
			fmt.Println("The dealer wins")
			os.Exit(0)
		}
	}
}
