package controllers

import (
	"Pruebas_D_GO/src/internal/model"
	"fmt"
)

type CrupierController struct {
	hand     []model.Card
	deckCtrl *DeckController
	handCtrl *HandController
}

func NewCrupierController(deckCtrl *DeckController, handCtrl *HandController) *CrupierController {
	return &CrupierController{
		hand:     []model.Card{},
		deckCtrl: deckCtrl,
		handCtrl: handCtrl,
	}
}

// Función para que el crupier tome cartas hasta llegar a un score de 17 o más
func (c *CrupierController) CuprierTurn() []model.Card {
	c.handCtrl.hand = c.hand
	dealerScore := c.handCtrl.HandValue()
	fmt.Println("\nScore del crupier antes de destapar la carta volteada:", dealerScore)
	for dealerScore <= 17 {
		card := c.deckCtrl.Deal()
		c.AddCard(card) // Agregar la carta a la mano del crupier
		c.handCtrl.hand = c.hand
		dealerScore = c.handCtrl.HandValue() // Actualizar la puntuación
		fmt.Println("\nEl crupier toma una carta.")
		fmt.Println("\nLa nueva carta del crupier es:", card.Value, "of", card.Suit) // Imprimir la nueva carta
	}
	return c.hand
}
func (c *CrupierController) ClearHand() {
	c.hand = []model.Card{}
}

// Función para añadir una carta a la mano
func (c *CrupierController) AddCard(card model.Card) {
	c.hand = append(c.hand, card)
}

// Funcion para mostrar la mano
func (c *CrupierController) ShowHand() string {
	var handString string
	for _, card := range c.hand {
		handString += card.Value + " of " + card.Suit + ", "
	}
	return handString[:len(handString)-2] // Eliminar la última coma y espacio
}

// Funcion para mostrar solo una carta
func (c *CrupierController) ShowHandHidden() string {
	if len(c.hand) == 0 {
		return ""
	}
	var handString string
	handString = "Hidden, " // Ocultar la primera carta
	for i, card := range c.hand {
		if i == 0 {
			continue
		}
		handString += card.Value + " of " + card.Suit + ", "
	}
	return handString[:len(handString)-2] // Eliminar la última coma y espacio
}
func (c *CrupierController) HandValue() int {
	c.handCtrl.hand = c.hand
	crupierScore := c.handCtrl.HandValue()
	return crupierScore
}
