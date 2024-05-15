package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

type Deck []Card

func NewDeck() Deck {
	var deck Deck
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			card := Card{Suit: suit, Value: value}
			deck = append(deck, card)
		}
	}
	return deck
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d *Deck) Deal() Card {
	card, remainingDeck := (*d)[0], (*d)[1:]
	*d = remainingDeck
	return card
}

// Función para que el jugador "hit" (tome más cartas)
func playerHit(playerHand []Card, deck Deck) []Card {
	var numEntrada int
	playerScore := handValue(playerHand)
	println("Score del jugador antes de dar un HIT", playerScore)
	fmt.Println("La mano del jugador antes de un HIT", playerHand)
	playerHand = append(playerHand, deck.Deal())
	fmt.Println("\nTu nueva carta es:", playerHand[len(playerHand)-1].Value, "of", playerHand[len(playerHand)-1].Suit)
	numEntrada++
	playerScore = handValue(playerHand)
	fmt.Println("Score del jugador despues de mostrar las nuevas cartas:", playerScore)
	fmt.Println("Numero de entrada", numEntrada)
	fmt.Println("La mano del jugador", playerHand)
	return playerHand
}

// Función para que el jugador "stand" (no tome más cartas)
func playerStand() {
	fmt.Println("\nTe paras.")
}

// Función para que el jugador "Double" (tome una carta mas y se para)
func playerDouble(playerHand []Card, deck Deck) int {
	playerHand = append(playerHand, deck.Deal())
	fmt.Println("\nTu nueva carta es:", playerHand[len(playerHand)-1].Value, "of", playerHand[len(playerHand)-1].Suit)
	return handValue(playerHand)
}

func CuprierTurn(dealerHand []Card, deck Deck) int {
	dealerScore := handValue(dealerHand)
	for dealerScore <= 17 {
		dealerHand = append(dealerHand, deck.Deal())
		dealerScore = handValue(dealerHand)
		fmt.Println("Score del crupier:", dealerScore)
		fmt.Println("\nEl crupier toma una carta.")
		fmt.Println("La nueva carta del crupier es:", dealerHand[len(dealerHand)-1].Value, "of", dealerHand[len(dealerHand)-1].Suit)
	}
	fmt.Println("Score del crupier:", dealerScore)
	return dealerScore
}

// Función para calcular el valor total de una mano de cartas
func handValue(hand []Card) int {
	total := 0
	numAces := 0
	for _, card := range hand {
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

func main() {
	mascartas := true
	playerBJ := false
	playerTurn := true
	turnNumber := 0
	var choice string
	deck := NewDeck()
	deck.Shuffle()
	PlayerScore := 0
	CrupierScore := 0

	// Ejemplo de reparto de cartas para el jugador y el crupier
	playerHand := []Card{deck.Deal(), deck.Deal()}
	dealerHand := []Card{deck.Deal(), deck.Deal()}

	fmt.Println("Cartas del jugador:")
	for _, card := range playerHand {
		fmt.Println(card.Value, "of", card.Suit)
	}

	PlayerScore = handValue(playerHand)
	fmt.Println("Score del jugador:", PlayerScore)
	fmt.Println("\nCartas del crupier:")
	fmt.Println("Hidden card")
	for i := 1; i < len(dealerHand); i++ {
		fmt.Println(dealerHand[i].Value, "of", dealerHand[i].Suit)
	}

	if PlayerScore == 21 {
		fmt.Println("Blackjack")
		playerTurn = false
	} else {
		for mascartas {
			fmt.Println("Score del jugador al entrar al for mascartas:", PlayerScore)
			fmt.Println("Mano del jugador al entrar al for mascartas:", playerHand)
			if PlayerScore < 21 {
				fmt.Println("Score del jugador al entrar al if player < 21 ", PlayerScore)
				fmt.Println("Mano del jugador al entrar al if player < 21 ", playerHand)
				if turnNumber == 0 {
					fmt.Printf("\n¿Quieres otra carta, te quedas o Doble? (H/S/D): ")
				} else {
					fmt.Printf("\n¿Quieres otra carta o te quedas? (H/S): ")
				}
				fmt.Scan(&choice)

				if choice == "H" {
					fmt.Println("Deckplayer old  del jugador", playerHand,  "en el turno ", turnNumber)
					turnNumber++
					PlayerScore = handValue(playerHand)
					playerHand = playerHit(playerHand, deck)
					fmt.Println("Score del jugador", PlayerScore,  "en el turno ", turnNumber)
					fmt.Println("deckplayer new del jugador", playerHand,  "en el turno ", turnNumber)
					PlayerScore = handValue(playerHand)
					fmt.Println("Score del jugador", PlayerScore,  "en el turno ", turnNumber)
				}
				if choice == "S" {
					playerStand()
					playerTurn = false
					mascartas = false
				} else if choice == "D" {
					PlayerScore = playerDouble(playerHand, deck)
					playerStand()
					fmt.Println("Score del jugador:", PlayerScore)
					playerTurn = false
					mascartas = false
				}
			}else if PlayerScore == 21 {
				break
			}else if PlayerScore > 21 {
				mascartas = false
				playerTurn = false
				break
			}

		}
		if PlayerScore == 21 || PlayerScore > 21 {
			if PlayerScore == 21 {
				fmt.Println("Blackjack")
				playerStand()
				playerTurn = false

			} else if PlayerScore > 21 {
				fmt.Println("Te pasaste de 21, Perdiste.")
				playerTurn = false
			}
		}else if PlayerScore > 21{
			fmt.Println("Te pasaste de 21, Perdiste.")
			playerTurn = false
		}
		if !playerTurn {
			fmt.Println("\nCartas del crupier:")
			for _, card := range dealerHand {
				fmt.Println(card.Value, "of", card.Suit)
			}
			CrupierScore = CuprierTurn(dealerHand, deck)
			if PlayerScore > CrupierScore && PlayerScore <= 21 {
				fmt.Println("Ganaste")
			} else if playerBJ && CrupierScore != 21 {
				fmt.Println("Ganaste con Blackjack")
			} else if PlayerScore < CrupierScore && CrupierScore <= 21 {
				fmt.Println("Perdiste")
			} else if PlayerScore == CrupierScore {
				fmt.Println("Empate")
			}
		}
	}

}
