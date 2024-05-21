package cmd

import (
	"model" 
)

type app struct {
	deck        model.Deck
	playerHand  []model.Card
	dealerHand  []model.Card
	playerScore int
	dealerScore int
}

func (a *app) Run() {
}

type App interface {
	Run()
}

func NewApp() App {
	a := &app{}
	a.deck.Shuffle()
	a.playerHand = []model.Card{}
	a.playerScore = 0
	a.dealerScore = 0

	return a
}
