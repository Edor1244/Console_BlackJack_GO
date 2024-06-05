package controllers

import (
	"Pruebas_D_GO/src/internal/model"
	"fmt"
)

type PlayerController struct {
	hand     HandInterface
	deckCtrl *DeckController
}

func NewPlayerController(deckCtrl *DeckController, handCtrl *HandController) *PlayerController {
	return &PlayerController{
		deckCtrl: deckCtrl,
		hand:     handCtrl,
	}
}
func (p *PlayerController) PlayerHit() int {
	card := p.deckCtrl.Deal()
	p.hand.AddCard(card)
	fmt.Println("\nTu nueva carta es:", card.Value, "of", card.Suit)
	newScore := p.hand.HandValue() // Llamar a handValue del handController
	fmt.Println("\nScore del jugador:", newScore)
	return newScore
}

// Función para que el jugador "stand" (no tome más cartas)
func PlayerStand() {
	fmt.Println("\nTe paras.")
}

func (p *PlayerController) HandValue() int {
	playerScore := p.hand.HandValue()
	return playerScore
}

// Función para que el jugador "Double" (tome una carta mas y se para)
func (p *PlayerController) PlayerDouble() int {
	card := p.deckCtrl.Deal()
	p.hand.AddCard(card)
	fmt.Println("\nTu nueva carta es:", card.Value, "of", card.Suit)
	newScore := p.hand.HandValue() // Llamar a handValue del handController
	return newScore
}

func (p *PlayerController) ShowHand() string {
	handString := p.hand.ShowHand()
	return handString // Eliminar la última coma y espacio
}

// Función para añadir una carta a la mano
func (p *PlayerController) AddCard(card model.Card) {
	p.hand.AddCard(card)
}

// Función para añadir una carta a la mano
func (p *PlayerController) ItsBlackJack() bool {
	var respuesta = p.hand.ItsBlackJack()
	return respuesta
}

func (p *PlayerController) ClearHand() {
	p.hand.ClearHand()
}
