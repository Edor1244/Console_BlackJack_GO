package controllers

import (
	"Pruebas_D_GO/src/internal/model"
	"fmt"
)

type playerController struct {
	hand []model.Card
	deckCtrl *deckController
	handCtrl *handController
}

func newPlayerController(deckCtrl *deckController, handCtrl *handController) *playerController {
	return &playerController{
		hand: model.Deck{},
		deckCtrl: deckCtrl,
		handCtrl: handCtrl,
	}
}
func (p *playerController) playerHit() int {
	card := p.deckCtrl.Deal()
	p.handCtrl.addCard(card)
	fmt.Println("\nTu nueva carta es:", p.hand[len(p.hand)-1].Value, "of", p.hand[len(p.hand)-1].Suit)

	newScore := p.handCtrl.handValue(p.hand) // Llamar a handValue del handController
	fmt.Println("\nScore del jugador:", newScore)
	return newScore
}

// Función para que el jugador "stand" (no tome más cartas)
func playerStand() {
	fmt.Println("\nTe paras.")
}

// Función para que el jugador "Double" (tome una carta mas y se para)
func(p *playerController) playerDouble() int {
	card := p.deckCtrl.Deal()
	p.handCtrl.addCard(card)
	fmt.Println("\nTu nueva carta es:", p.hand[len(p.hand)-1].Value, "of", p.hand[len(p.hand)-1].Suit)
	newScore := p.handCtrl.handValue(p.hand) // Llamar a handValue del handController
	return newScore
}


