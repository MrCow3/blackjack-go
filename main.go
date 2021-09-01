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

var finalCount = []int{playerAmount, dealerAmount}

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
			} else if strings.Compare("stand", text) == 0 {
				if 21 >= finalCount[0] {
					x := 21 - finalCount[0]
					y := 21 - finalCount[1]
					if x > y {
						fmt.Println("You win")
						os.Exit(0)
					}
					if y > x {
						fmt.Println("The dealer wins")
						os.Exit(0)
					}
				}
			}
		} else if playerAmount > 21 {
			fmt.Println("Bust!")
			break
		} else if playerAmount == 21 {
			fmt.Println("Blackjack!")
			break
		}
	}
}

func randCard(n int) int {
	rand.Seed(time.Now().UnixNano())
	in := cards
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	playerAmount = playerAmount + pick
	finalCount[0] = playerAmount
	return 0
}

func randCardDealer(n int) int {
	rand.Seed(time.Now().UnixNano())
	in := cards
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	dealerAmount = dealerAmount + pick
	finalCount[1] = dealerAmount
	return 0
}
