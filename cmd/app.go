package cmd

import (
	"Pruebas_D_GO/src/game/controllers"
	"fmt"
)

type app struct {
	deckCtrl    *controllers.DeckController
	playerCtrl  *controllers.PlayerController
	crupierCtrl *controllers.CrupierController
	handCtrl    *controllers.HandController
}

func (a *app) Run() {
	// ... (Lógica del juego aquí)
	playAgain := true
	var choice string
	for playAgain {
		a.LimpiarDatos()
		fmt.Printf("\n ¿Quieres iniciar una Partida? S/N ")
		fmt.Scan(&choice)
		switch choice {
		case "S":
			fmt.Println("Ok Juguemos")
			a.iniciarJuego()
		case "N":
			fmt.Println("Ok, Bye")
			playAgain = false
		}
	}
}

type App interface {
	Run()
}

func NewApp() App {
	a := &app{}
	a.deckCtrl = controllers.NewDeckController()
	a.handCtrl = controllers.NewHandController()
	a.playerCtrl = controllers.NewPlayerController(a.deckCtrl, a.handCtrl)
	a.crupierCtrl = controllers.NewCrupierController(a.deckCtrl, a.handCtrl)
	return a
}

func (a *app) iniciarJuego() {
	fmt.Println("Entro a inciar Juego")

	a.deckCtrl.Shuffle() // Barajar el mazo

	// Repartir cartas iniciales
	for i := 0; i < 2; i++ {
		a.playerCtrl.AddCard(a.deckCtrl.Deal())
		a.crupierCtrl.AddCard(a.deckCtrl.Deal())
	}

	// Calcular puntuaciones iniciales
	playerScore := a.playerCtrl.HandValue()

	fmt.Println("Cartas del jugador:")
	fmt.Println(a.playerCtrl.ShowHand())
	fmt.Println("Score del jugador:", playerScore)

	// Verificar si hay Blackjack (opcional, si tienes el método isBlackjack)
	if a.playerCtrl.ItsBlackJack() {
		fmt.Println("¡Blackjack!")
		a.finalizarJuegoBJP()
	} else if a.playerCtrl.HandValue() > 21 {
		a.finalizarJuegoExceso()
	} else {
		a.continuarJuego()
	}
}

func (a *app) finalizarJuegoBJP() {
	fmt.Println("\n--- Resultados Finales ---")
	fmt.Println("Mano del jugador:", a.playerCtrl.ShowHand(), "- Puntuación:", a.playerCtrl.HandValue())
	fmt.Println("\n¡El jugador gana con Blackjack!") // Mensaje específico para este caso
}
func (a *app) finalizarJuegoExceso() {
	fmt.Println("\n--- Resultados Finales ---")
	fmt.Println("Mano del jugador:", a.playerCtrl.ShowHand(), "- Puntuación:", a.playerCtrl.HandValue())
	fmt.Println("\n¡El jugador se paso de 21 Pierdes!") // Mensaje específico para este caso
}
func (a *app) continuarJuego() {

	fmt.Println("Cartas del crupier:")
	fmt.Println(a.crupierCtrl.ShowHandHidden()) // Mostrar solo la segunda carta
	var choice string
	playerTurn := true
	gameInProcess := true
	turns := 0

	for playerTurn && gameInProcess {
		playerScore := a.playerCtrl.HandValue()
		if playerScore > 21 {
			gameInProcess = false
		} else {
			if turns == 0 {
				fmt.Printf("\n¿Quieres otra carta o te quedas? (H/S/D): ")
				fmt.Scan(&choice)
			} else {
				fmt.Printf("\n¿Quieres otra carta o te quedas? (H/S): ")
				fmt.Scan(&choice)
			}
			turns++
			switch choice {
			case "H":
				a.playerCtrl.PlayerHit()
				playerScore := a.playerCtrl.HandValue()
				fmt.Println("Cartas del jugador:")
				fmt.Println(a.playerCtrl.ShowHand())
				fmt.Println("Score del jugador:", playerScore)

				if playerScore == 21 {
					fmt.Println("¡Tienes 21!")
					playerTurn = false
				}
			case "S":
				fmt.Println("Te quedas.")
				playerTurn = false
			case "D":
				a.playerCtrl.PlayerDouble()
				playerScore := a.playerCtrl.HandValue()
				if playerScore > 21 {
					gameInProcess = false
				} else {
					fmt.Println("Cartas del jugador:")
					fmt.Println(a.playerCtrl.ShowHand())
					fmt.Println("Score del jugador:", playerScore)
					playerTurn = false
				}
			default:
				fmt.Println("Opción no válida. Intenta de nuevo.")
			}
		}
	}
	if !playerTurn && gameInProcess {
		a.jugarTurnoCrupier()
	} else if !gameInProcess {
		a.finalizarJuegoExceso()
	}
}

func (a *app) jugarTurnoCrupier() {
	fmt.Println("\nEs el turno del crupier.")
	fmt.Println("\nCartas del crupier:")
	fmt.Println(a.crupierCtrl.ShowHand())
	a.crupierCtrl.CuprierTurn()
	crupierScore := a.crupierCtrl.HandValue()

	fmt.Println("\nCartasdel crupier: ", a.crupierCtrl.ShowHand())
	fmt.Println("\nScore del crupier:", crupierScore)

	playerScore := a.playerCtrl.HandValue()

	if crupierScore > 21 {
		fmt.Println("¡Ganaste! El crupier se pasó de 21.")
	} else if playerScore > crupierScore {
		fmt.Println("¡Ganaste!")
	} else if playerScore < crupierScore {
		fmt.Println("¡Perdiste!")
	} else {
		fmt.Println("¡Empate!")
	}
}
func (a *app) LimpiarDatos() {
	a.handCtrl.ClearHand()
	a.crupierCtrl.ClearHand()
	a.playerCtrl.ClearHand()
}
