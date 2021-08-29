package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var playerAmmount int = 0

var cards = [9]int{2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		if playerAmmount < 21 {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')

			text = strings.Replace(text, "\n", "", -1)
			fmt.Println("Would you like to hit")
			if strings.Compare("yes", text) == 0 {
				randCard(0)
				fmt.Println(playerAmmount)
			}
		} else if playerAmmount > 21 {
			fmt.Println("Bust!")
			os.Exit(0)
		}
	}
}

func randCard(n int) int {
	rand.Seed(time.Now().UnixNano())
	in := cards
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	playerAmmount = playerAmmount + pick
	return 0
}
