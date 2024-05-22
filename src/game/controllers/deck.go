package controllers

import (
	"Pruebas_D_GO/src/model"
	"math/rand"
	"time"
)

type DeckController struct {
	deck model.Deck
}

func NewDeckController() *DeckController {
    d := &DeckController{
        deck: NewDeck(), // Llama al método NewDeck para crear y asignar la baraja
    }
    d.Shuffle()
    return d
}

//Estructura de la mano
func NewDeck() model.Deck {
	var deck model.Deck
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			card := model.Card{Suit: suit, Value: value}
			deck = append(deck, card)
		}
	}
	return deck
}

//Función para barajar las cartas
func (d *DeckController) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d.deck {
		newPosition := r.Intn(len(d.deck) - 1)
		d.deck[i], d.deck[newPosition] = d.deck[newPosition], d.deck[i]
	}
}

//Función para repartir las cartas
func (d * DeckController) Deal() model.Card {
	card, remainingDeck := d.deck[0], d.deck[1:]
	d.deck = remainingDeck
	return card
}



