package main

import "fmt"

var playerAmmount int = 0

var cards = map[string]int{
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
}

func main() {
	fmt.Println(cards["three"])
}
