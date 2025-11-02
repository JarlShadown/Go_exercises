package main

func main() {
	cards := loadFromFile("deck.csv")
	cards.print()

	//cards := newDeck()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//fmt.Println("----------------")
	//remainingCards.print()
	//saveFile(cards)
}
