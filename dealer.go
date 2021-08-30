package main

import "fmt"

var dealerAmount int = 0

func main() {
	fmt.Println(dealerAmount)
	for {
		if dealerAmount < 21 {
			randCard(0)
		}
		
		if dealerAmount > 21 {
			fmt.Println("Dealer busted!")
		}
		
		if dealerAmount = 21 {
			fmt.Println("Dealer Blackjack!")
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
