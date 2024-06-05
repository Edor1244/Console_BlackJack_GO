package controllers

import (
	"Pruebas_D_GO/src/internal/model"
	"fmt"
)

type PlayerController struct {
	hand     []model.Card
	deckCtrl *DeckController
	handCtrl *HandController
}

func NewPlayerController(deckCtrl *DeckController, handCtrl *HandController) *PlayerController {
	return &PlayerController{
		hand:     model.Deck{},
		deckCtrl: deckCtrl,
		handCtrl: handCtrl,
	}
}
func (p *PlayerController) PlayerHit() int {
	card := p.deckCtrl.Deal()
	p.AddCard(card)
	fmt.Println("\nTu nueva carta es:", p.hand[len(p.hand)-1].Value, "of", p.hand[len(p.hand)-1].Suit)
	p.handCtrl.hand = p.hand
	newScore := p.handCtrl.HandValue() // Llamar a handValue del handController
	fmt.Println("\nScore del jugador:", newScore)
	return newScore
}

// Función para que el jugador "stand" (no tome más cartas)
func PlayerStand() {
	fmt.Println("\nTe paras.")
}

func (p *PlayerController) HandValue() int {
	p.handCtrl.hand = p.hand
	playerScore := p.handCtrl.HandValue()
	return playerScore
}

// Función para que el jugador "Double" (tome una carta mas y se para)
func (p *PlayerController) PlayerDouble() int {
	card := p.deckCtrl.Deal()
	p.AddCard(card)
	fmt.Println("\nTu nueva carta es:", p.hand[len(p.hand)-1].Value, "of", p.hand[len(p.hand)-1].Suit)
	newScore := p.handCtrl.HandValue() // Llamar a handValue del handController
	return newScore
}

func (p *PlayerController) ShowHand() string {
	var handString string
	for _, card := range p.hand {
		handString += card.Value + " of " + card.Suit + ", "
	}
	return handString[:len(handString)-2] // Eliminar la última coma y espacio
}

// Función para añadir una carta a la mano
func (p *PlayerController) AddCard(card model.Card) {
	p.hand = append(p.hand, card)
}

// Función para añadir una carta a la mano
func (p *PlayerController) ItsBlackJack() bool {
	var respuesta = p.handCtrl.ItsBlackJack()
	return respuesta
}

func (p *PlayerController) ClearHand() {
	p.hand = []model.Card{}
}
