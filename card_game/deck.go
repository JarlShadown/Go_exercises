package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck
// which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	//then make it randown

	return cards.shuffle()
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() deck {
	//use the rand package
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func saveFile(d deck) {
	err := os.WriteFile("deck.csv", []byte(d.toString()), 0666)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}
func loadFromFile(route string) deck {
	file, err := os.ReadFile(route)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return deck(strings.Split(string(file), ","))
}
