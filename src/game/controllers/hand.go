package controllers

import "Pruebas_D_GO/src/model"

type handController struct {
	hand []model.Card	
}

func newHandController() *handController {
    return &handController{
        hand: []model.Card{}, // Inicializar la mano vacía
    }
}

// Función para añadir una carta a la mano
func (h *handController) addCard(card model.Card) {
    h.hand = append(h.hand, card)
}

//Funcion para ver si una carta es BlackJack
func(h *handController) itsBlackJack() bool{
	total := 0;
	total = h.handValue()
	if(total == 21){
		return true
	}else{
		return false
	}
}

// Función para calcular el valor total de una mano de cartas
func (h *handController) handValue() int {
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
