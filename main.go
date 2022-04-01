package main

import "fmt"

func main() {
	i := 0
	fileName := "my_cards"
	fmt.Print("Enter number: ")
	fmt.Scanf("%v", &i)
	if i == 1 {
		// create new deck and shuffle
		cards := newDeck()
		cards.shuffle()
		fmt.Println("=== Shuffled ===")
		cards.print()

		// deal cards
		hand, remainingCards := deal(cards, 5)
		fmt.Println("=== On Hand ===")
		hand.print()
		fmt.Println("=== Remaining ===")
		remainingCards.print()

		//save to file
		cards.saveToFile(fileName)
		fmt.Printf("*** saved, %v ***\n", fileName)
		// cards = append(cards, hand...)

	} else if i == 2 {
		// read deck from file
		cards := newDeckFromFile(fileName)
		cards.print()
	}

}
