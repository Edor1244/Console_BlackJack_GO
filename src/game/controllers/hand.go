package controllers

import (
	"Pruebas_D_GO/src/internal/model"
)

type HandController struct {
	hand []model.Card
}

func NewHandController() *HandController {
	return &HandController{
		hand: []model.Card{}, // Inicializar la mano vacía
	}
}

// Funcion para limpiar manos
func (h *HandController) ClearHand() {
	h.hand = []model.Card{}
}

func (h *HandController) ShowHand() string {
	var handString string
	for _, card := range h.hand {
		handString += card.Value + " of " + card.Suit + ", "
	}
	return handString[:len(handString)-2] // Eliminar la última coma y espacio
}

// Funcion para mostrar solo una carta cuprier
func (h *HandController) ShowHandHidden() string {
	if len(h.hand) == 0 {
		return ""
	}
	var handString string
	handString = "Hidden, " // Ocultar la primera carta
	for i, card := range h.hand {
		if i == 0 {
			continue
		}
		handString += card.Value + " of " + card.Suit + ", "
	}
	return handString[:len(handString)-2] // Eliminar la última coma y espacio
}

// Función para añadir una carta a la mano
func (h *HandController) AddCard(card model.Card) {
	h.hand = append(h.hand, card)
}

// Funcion para ver si una carta es BlackJack
func (h *HandController) ItsBlackJack() bool {
	total := 0
	total = h.HandValue()
	if total == 21 {
		return true
	} else {
		return false
	}
}

// Función para calcular el valor total de una mano de cartas
func (h *HandController) HandValue() int {
	total := 0
	numAces := 0
	for _, card := range h.hand {
		switch card.Value {
		case "Ace":
			total += 11
			numAces++
		case "Two":
			total += 2
		case "Three":
			total += 3
		case "Four":
			total += 4
		case "Five":
			total += 5
		case "Six":
			total += 6
		case "Seven":
			total += 7
		case "Eight":
			total += 8
		case "Nine":
			total += 9
		default:
			total += 10
		}
	}
	// Ajustar el valor de los ases si el total excede 21
	for numAces > 0 && total > 21 {
		total -= 10
		numAces--
	}
	return total
}
