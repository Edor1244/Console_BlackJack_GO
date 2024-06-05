package controllers

import "Pruebas_D_GO/src/internal/model"

type HandInterface interface {
	ClearHand()
	AddCard(card model.Card)
	HandValue() int
	ItsBlackJack() bool
	ShowHand() string
	ShowHandHidden() string
}

type Player struct {
	hand HandInterface
	// otros campos...
}

type Cuprier struct {
	hand HandInterface
	// otros campos...
}
