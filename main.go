package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var playerAmount int = 0
var dealerAmount int = 0

var cards = [9]int{2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {

	dealer()

	for {
		if playerAmount < 21 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Hit or stand")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')

			text = strings.Replace(text, "\n", "", -1)
			if strings.Compare("hit", text) == 0 {
				randCard(0)
				fmt.Println(playerAmount)
			}
		} else if playerAmount > 21 {
			fmt.Println("Bust!")
			os.Exit(0)
		} else if playerAmount == 21 {
			fmt.Println("Blackjack!")
			os.Exit(0)
		}
	}
}

func randCard(n int) int {
	rand.Seed(time.Now().UnixNano())
	in := cards
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	playerAmount = playerAmount + pick
	return 0
}

func randCardDealer(n int) int {
	rand.Seed(time.Now().UnixNano())
	in := cards
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	dealerAmount = dealerAmount + pick
	return 0
}
