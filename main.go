package main

import (
	"Pruebas_D_GO/cmd"
	"fmt"
)

func main() {

	app := cmd.NewApp()
	app.Run()

	mascartas := true
	playerBJ := false
	playerTurn := true
	var crupierTurn bool
	turnNumber := 0
	var choice string


	if PlayerScore == 21 {
		fmt.Println("Ganaste con Blackjack")
		playerBJ = true
		playerTurn = false
	} else {
		fmt.Println("\nCartas del crupier:")
		fmt.Println("Hidden card")
		for i := 1; i < len(dealerHand); i++ {
			fmt.Println(dealerHand[i].Value, "of", dealerHand[i].Suit)
		}
		for mascartas {
			if PlayerScore < 21 {
				if turnNumber == 0 {
					fmt.Printf("\n¿Quieres otra carta, te quedas o Doble? (H/S/D): ")
				} else {
					fmt.Printf("\n¿Quieres otra carta o te quedas? (H/S): ")
				}
				fmt.Scan(&choice)

				if choice == "H" {
					turnNumber++
					PlayerScore = handValue(playerHand)
					playerHand = playerHit(playerHand, deck)
					PlayerScore = handValue(playerHand)
				}
				if choice == "S" {
					playerStand()
					playerTurn = false
					mascartas = false
					crupierTurn = true
					break
				} else if choice == "D" {
					PlayerScore = playerDouble(playerHand, deck)
					playerStand()
					fmt.Println("\nScore del jugador:", PlayerScore)
					playerTurn = false
					mascartas = false
					crupierTurn = true
				}
			} else if PlayerScore == 21 {
				fmt.Println("Obtuviste 21, te quedas.")
				CrupierScore = CuprierTurn(dealerHand, deck)
				if CrupierScore != 21 {
					fmt.Println("\n¡Ganaste con 21!")
					crupierTurn = false
					playerTurn = false
					break
				} else {
					fmt.Println("\n¡Empate!")
					crupierTurn = false
					playerTurn = false
					break
				}
			} else if PlayerScore > 21 {
				fmt.Println("\nTe pasaste de 21, Perdiste.")
				crupierTurn = false
				playerTurn = false
				break
			}
		}
		if crupierTurn {
			if !playerTurn {
				fmt.Println("\nEs el turno del crupier.")
				fmt.Println("\nCartas del crupier:")
				for _, card := range dealerHand {
					fmt.Println("\n", card.Value, "of", card.Suit)
				}
				CrupierScore = CuprierTurn(dealerHand, deck)
				if CrupierScore > 21 {
					fmt.Println("\n¡Ganaste! El crupier se pasó de 21.")
				} else if PlayerScore > CrupierScore && PlayerScore <= 21 {
					fmt.Println("\n¡Ganaste!")
				} else if playerBJ && CrupierScore != 21 {
					fmt.Println("\n¡Ganaste con Blackjack!")
				} else if PlayerScore < CrupierScore && CrupierScore <= 21 {
					fmt.Println("\n¡Perdiste!")
				} else if PlayerScore == CrupierScore {
					fmt.Println("\n¡Empate!")
				} else {
					fmt.Println("\n¡Perdiste! El crupier gana.")
				}
			}
			fmt.Println("\n Acabo el juego.")
		} else {
			fmt.Println("\nAcabo el juego.")
		}
	}
}
