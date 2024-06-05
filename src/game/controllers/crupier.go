package controllers

import (
	"Pruebas_D_GO/src/internal/model"
	"fmt"
)

type CrupierController struct {
	hand     HandInterface
	deckCtrl *DeckController
}

func NewCrupierController(deckCtrl *DeckController, handCtrl *HandController) *CrupierController {
	return &CrupierController{
		deckCtrl: deckCtrl,
		hand:     handCtrl,
	}
}

// Función para que el crupier tome cartas hasta llegar a un score de 17 o más
func (c *CrupierController) CuprierTurn() int {
	dealerScore := c.hand.HandValue()
	fmt.Println("\nScore del crupier antes de destapar la carta volteada:", dealerScore)
	for dealerScore <= 17 {
		card := c.deckCtrl.Deal()
		c.AddCard(card)                  // Agregar la carta a la mano del crupier
		dealerScore = c.hand.HandValue() // Actualizar la puntuación
		fmt.Println("\nEl crupier toma una carta.")
		fmt.Println("\nLa nueva carta del crupier es:", card.Value, "of", card.Suit) // Imprimir la nueva carta
	}
	return c.hand.HandValue()
}

func (c *CrupierController) ClearHand() {
	c.hand.ClearHand()
}

// Función para añadir una carta a la mano
func (c *CrupierController) AddCard(card model.Card) {
	c.hand.AddCard(card)
}

// Funcion para mostrar la mano
func (c *CrupierController) ShowHand() string {
	return c.hand.ShowHand() // Eliminar la última coma y espacio
}

// Funcion para mostrar solo una carta
func (c *CrupierController) ShowHandHidden() string {
	handString := c.hand.ShowHandHidden() // Eliminar la última coma y espacio
	return handString
}
func (c *CrupierController) HandValue() int {
	crupierScore := c.hand.HandValue()
	return crupierScore
}
