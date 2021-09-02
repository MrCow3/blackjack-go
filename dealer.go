package main

import (
	"fmt"
	"os"
	"time"
)

func dealer() {
	for {
		if dealerAmount < 17 {
			fmt.Println("The dealer hits")
			randCardDealer(13)
			time.Sleep(1 * time.Second)
		} else if dealerAmount >= 17 {
			fmt.Println("The dealer stands")
			break
		} else if dealerAmount > 21 {
			fmt.Println("Dealer busted!")
			fmt.Println("You win")
			os.Exit(0)
		} else if dealerAmount == 21 {
			fmt.Println("Dealer Blackjack!")
			fmt.Println("The dealer wins")
			os.Exit(0)
		}
	}
}
