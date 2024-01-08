package main

//You got it! We'll figure out a way to deal with the unused 'index' variable in just a minute.

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//cards := newDeck()
	//fmt.Println(cards.saveToFile("my_file"))
	//fmt.Println(cards.toString())
	// hand, remaingCards := deal(cards, 5)
	// remaingCards.print()
	// hand.print()
}
