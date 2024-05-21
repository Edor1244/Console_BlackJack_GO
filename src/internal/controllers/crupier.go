package controllers

import (
	"Pruebas_D_GO/src/internal/model"
	"fmt"
)

type crupierController struct {
	hand []model.Card
	deckCtrl *deckController
	handCtrl *handController
}

func newCrupierController(deckCtrl *deckController, handCtrl *handController) *crupierController {
	return &crupierController{
		hand: []model.Card{},
		deckCtrl: deckCtrl,
		handCtrl: handCtrl,
	}
}
// Función para que el crupier tome cartas hasta llegar a un score de 17 o más
func (c *crupierController) CuprierTurn() []model.Card {
    dealerScore := c.handCtrl.handValue(c.hand)
    fmt.Println("Score del crupier al destapar carta volteada:", dealerScore)
    for dealerScore <= 17 {
        card := c.deckCtrl.Deal()
		c.handCtrl.addCard(card)// Agregar la carta a la mano del crupier
        dealerScore = c.handCtrl.handValue(c.hand) // Actualizar la puntuación
        fmt.Println("\nEl crupier toma una carta.")
        fmt.Println("\nLa nueva carta del crupier es:", card.Value, "of", card.Suit) // Imprimir la nueva carta
        fmt.Println("\nScore del crupier:", dealerScore)
    }
    return c.hand
}

