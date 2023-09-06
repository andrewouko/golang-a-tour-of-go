package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PlayingCard struct {
	Suit string
	Rank string
}

func NewPlayingCard(suit string, rank string) *PlayingCard {
	return &PlayingCard{suit, rank}
}

func (card *PlayingCard) String() string {
	return fmt.Sprintf("%v of %v", card.Rank, card.Suit)
}

// Add card interface to ensure
// cards in Deck are similar and comparable
type Card interface {
	fmt.Stringer
	Name() string
}

// create generic C
// to handle cards of different types
type Deck[C Card] struct {
	cards []C
}

// use generic Deck with generic C
// for any card type
func (deck *Deck[C]) AddCard(card C) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck[C]) String() string {
	return fmt.Sprintf("Deck %T , %v card(s) in the deck", deck, len(deck.cards))
}

// specify type of Deck generic
// to be of *PlayingCard type
// for use in the deck.cards slice
func NewPlayingCardDeck() *Deck[*PlayingCard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	var deck = &Deck[*PlayingCard]{}

	for _, suit := range suits {
		for _, rank := range ranks {
			// add card to deck using new card func
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

func (deck *Deck[C]) RandomCard() C {
	// random num between 0 and len of cards in deck
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	index := r.Intn(len(deck.cards))

	return deck.cards[index]
}

// ADD new type of card
type TradingCard struct {
	name string
}

func NewTradingCard(name string) *TradingCard {
	return &TradingCard{name}
}

func (tc *TradingCard) String() string {
	return fmt.Sprintf("Trading card name: %v", tc.name)
}

func NewTradingCardDeck() *Deck[*TradingCard] {
	cards := []string{"Sammy", "Droplets", "Spaces", "App Platform"}
	var deck = &Deck[*TradingCard]{}
	for _, name := range cards {
		deck.cards = append(deck.cards, NewTradingCard(name))
	}
	return deck
}

// Add name method to satisfy interface requirements
func (card *TradingCard) Name() string {
	return card.String()
}
func (card *PlayingCard) Name() string {
	return card.String()
}

func printRandomCardFromDeck[C Card](deck *Deck[C]) {
	card := deck.RandomCard()
	fmt.Println(deck)
	fmt.Println(card)
}

func main() {
	/* // create new playing deck
	deck := NewPlayingCardDeck()
	fmt.Println("New Playing Card Deck:", deck)

	// pick random card
	card := deck.RandomCard()
	fmt.Println("Random card picked", card) */

	// CODE REMOVED
	// No need for type check as we are using generics
	/* // check type of card
	normal_card, ok := card.(*Card)
	if !ok {
		fmt.Println("Card is not of type *Card")
		os.Exit(1)
	}
	fmt.Printf("Card suit: %v \nCard rank: %v", normal_card.Suit, normal_card.Rank)
	*/

	/* trading_cards_deck := NewTradingCardDeck()
	fmt.Println("New Trading Card Deck: ", trading_cards_deck)

	trading_card := trading_cards_deck.RandomCard()
	fmt.Println("Random card picked", trading_card) */
	playing_cards_deck := NewPlayingCardDeck()
	trading_cards_deck := NewTradingCardDeck()
	printRandomCardFromDeck(playing_cards_deck)
	printRandomCardFromDeck(trading_cards_deck)
}
